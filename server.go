package potato

import (
	"net/http"
	"os"
	"context"
	"log"
	"os/signal"
	"github.com/gin-gonic/gin"
	"github.com/flameous/PotatoPartyBackend/db"
	"github.com/flameous/PotatoPartyBackend/types"
	"strconv"
	"encoding/json"
	"time"
	"fmt"
)

// server holds config and gin server instance
type server struct {
	storage *db.Storage
	wr      *websocketRoom
	online  chan []int64
}

// NewServer creates server with taken config
func NewServer() *server {
	online := make(chan []int64, 100)
	return &server{wr: newWebsocketRoom(online), online: online}
}

// Serve serves perfectly!
func (s *server) Serve(exit chan bool) {
	go s.createEventForOnlineUsers()

	dbHost := os.Getenv("DB_HOST")
	if dbHost != "localhost" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	storage, err := db.ConnectToDatabase(dbHost)
	if err != nil {
		log.Fatalf("failed to connect to db - reason: %v", err)
	}
	s.storage = storage

	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)

	go func() {
		select {
		case <-exit:
		case <-sig:
		}

		log.Println("user interruption")
		err := srv.Shutdown(context.Background())
		if err != nil {
			log.Println("server shutdown", err)
		}
	}()

	// logger middleware
	engine.Use(func(c *gin.Context) {
		log.Printf("new request. %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	engine.POST("/auth", s.auth)
	engine.Use(s.authenticateMiddleware)

	engine.Any("/ws", func(c *gin.Context) {
		s.wr.addClient(c.MustGet("user").(*types.User), c.Writer, c.Request)
	})

	user := engine.Group("/user")
	user.GET("", s.getAllUsers)
	user.GET("/:id", s.getUserByID)
	user.GET("/:id/friends", s.getUserFriends)

	event := engine.Group("/event")
	event.GET("", s.getAllEvents)
	event.POST("", s.createNewEvent)

	event.GET("/:id", s.getEventByID)
	event.POST("/:id/attend", s.attendToEvent)
	event.POST("/:id/refuse", s.refuseAttendance)

	log.Println("server successfuly started")
	log.Println("server stopped", srv.ListenAndServe())
}

// auth by nickname and password
func (s *server) auth(c *gin.Context) {
	nickname, ok := c.GetPostForm("nickname")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, "missing 'id' field")
		return
	}
	pass, ok := c.GetPostForm("password")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, "missing 'password' field")
		return
	}

	log.Println("/auth", nickname, pass)
	u, err := s.storage.GetUserByNickname(nickname)
	if err != nil {
		log.Printf("get user. id = %s, error = %v", c.Param("id"), err)
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if u == nil {
		c.IndentedJSON(404, "not found")
		return
	}
	// check passwords
	if u.Password != pass {
		c.IndentedJSON(http.StatusBadRequest, "wrong password")
		return
	}

	// token generation
	u.GenerateNewToken()
	err = s.storage.UpdateUserToken(u)
	if err != nil {
		log.Println("update token error", err, u)
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(200, u)
}

func (s *server) authenticateMiddleware(c *gin.Context) {
	var token string
	if c.Request.Method == "GET" {
		token = c.Query("token")
	} else {
		token = c.PostForm("token")
	}

	u, err := s.storage.GetUserByToken(token)
	if err != nil {
		log.Println("check token error", err, token, u)
	}

	if u != nil {
		// store user data in context
		c.Set("user", u)
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			"need to authorize. POST /auth, fields 'nickname' and 'password'. passed token = "+token)
	}
}

func (s *server) getUserByID(c *gin.Context) {
	s.validateOrAbort(c)
	u, err := s.storage.GetUserByID(c.GetInt64("id"))
	if err != nil {
		log.Printf("get user. id = %s, error = %v", c.Param("id"), err)
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if u == nil {
		c.IndentedJSON(404, "not found")
		return
	}
	c.IndentedJSON(200, u)
}

func (s *server) getAllUsers(c *gin.Context) {
	users, err := s.storage.GetAllUsers()
	if err != nil {
		log.Println("get /user", err)
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(200, users)
}

func (s *server) getUserFriends(c *gin.Context) {
	c.IndentedJSON(200, "[1,2,3]")
}

func (s *server) getAllEvents(c *gin.Context) {
	events, err := s.storage.GetAllEvents(c.MustGet("user").(*types.User))
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(200, events)
}

func (s *server) createNewEvent(c *gin.Context) {
	data, ok := c.GetPostForm("event")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, "missing 'event' form-key")
		return
	}
	var event types.Event
	err := json.Unmarshal([]byte(data), &event)
	if err != nil {
		log.Println(err, data)
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	event.Owner = c.MustGet("user").(*types.User)
	id, err := s.storage.CreateNewEvent(&event)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	for _, v := range event.Attendees {
		s.wr.sendEventInvitationToClient(v.ID, &event)
	}
	c.IndentedJSON(200, id)
}

func (s *server) getEventByID(c *gin.Context) {
	s.validateOrAbort(c)
	e, err := s.storage.GetEventByID(c.GetInt64("id"), c.MustGet("user").(*types.User))
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	if e == nil {
		c.IndentedJSON(404, "not found")
		return
	}
	c.IndentedJSON(200, e)
}

func (s *server) attendToEvent(c *gin.Context) {
	s.validateOrAbort(c)
	err := s.storage.AttendToEvent(c.GetInt64("id"), s.getAuthUserID(c))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.String(200, "ok")
	}
}

func (s *server) refuseAttendance(c *gin.Context) {
	s.validateOrAbort(c)
	err := s.storage.RefuseAttendance(c.GetInt64("id"), s.getAuthUserID(c))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.String(200, "ok")
	}
}

func (s *server) validateOrAbort(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Println("parsed id = ", id)
	c.Set("id", id)
}

func (s *server) getAuthUserID(c *gin.Context) int64 {
	return c.MustGet("user").(*types.User).ID
}

func (s *server) createEventForOnlineUsers() {
	var num int = 0
	var online []int64
	for {
		online = <-s.online
		if len(online) < 2 {
			continue
		}
		num++

		users := make([]*types.User, 0)
		for _, id := range online {
			u, _ := s.storage.GetUserByID(id)
			users = append(users, u)
		}

		u, _ := s.storage.GetUserByNickname("Potato Party Bot")
		event := types.Event{
			Owner: u,
			Name:  fmt.Sprintf("Автоматическая тусовка номер %d !", num),
			Description: "Вы были выбраны высшим разумом, чтобы устроить тусовку." +
				" Наверное. Если не устроете, мы продадим ваши персональные данные :)",
			IsPrivate: false,
			Lat:       53.92667,
			Lon:       27.682518,
			Date:      time.Now(),
			Attendees: users,
			Interests: s.storage.RandomAllInterests(),
		}

		_, err := s.storage.CreateNewEvent(&event)
		if err != nil {
			log.Println(err)
		}

		for _, v := range event.Attendees {
			s.wr.sendEventInvitationToClient(v.ID, &event)
		}
	}
}
