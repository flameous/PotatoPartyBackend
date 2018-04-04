package potato

import (
	"log"
	"github.com/gorilla/websocket"
	"sync"
	"github.com/flameous/PotatoPartyBackend/types"
	"net/http"
	"time"
	"encoding/json"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	up = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}

var up websocket.Upgrader

type websocketRoom struct {
	Clients map[int64]*Client
	mux     *sync.RWMutex
	chat    chan *ChatMessage
}

type ChatMessage struct {
	Type string      `json:"type"`
	User *types.User `json:"user"`
	Text string      `json:"text"`
	Lat  float32     `json:"lat"`
	Lon  float32     `json:"lon"`
}

func newWebsocketRoom(out chan []int64) *websocketRoom {
	wr := websocketRoom{
		Clients: make(map[int64]*Client),
		mux:     &sync.RWMutex{},
		chat:    make(chan *ChatMessage),
	}

	go wr.cycle(out)
	go wr.chatCycle()

	return &wr
}

func (wr *websocketRoom) cycle(out chan []int64) {
	for {
		time.Sleep(time.Second * 30)
		clients := wr.readAllClients()
		out <- clients
	}
}

func (wr *websocketRoom) chatCycle() {
	for {
		msg := <-wr.chat
		ids := wr.readAllClients()

		for _, id := range ids {
			wr.sendMessageToChat(id, msg)
		}
	}
}

func (wr *websocketRoom) readAllClients() []int64 {
	wr.mux.RLock()
	defer wr.mux.RUnlock()

	var users []int64
	for k := range wr.Clients {
		users = append(users, k)
	}
	return users
}

func (wr *websocketRoom) addClient(u *types.User, w http.ResponseWriter, r *http.Request) {
	wr.mux.Lock()
	defer wr.mux.Unlock()

	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WEBSOCKET. CONNECT", err)
		return
	}

	c := &Client{u, conn}
	go c.handleIncomingMessage(wr.chat)
	wr.Clients[u.ID] = c
}

type Client struct {
	u    *types.User
	conn *websocket.Conn
}

func (c *Client) handleIncomingMessage(ch chan *ChatMessage) {
	for {
		_, data, err := c.conn.ReadMessage()
		log.Println(string(data), err)
		if err != nil {
			log.Println("WEBSOCKET. INCOMING MESSAGE", err)
			return
		}

		var chm ChatMessage
		if err = json.Unmarshal(data, &chm); err != nil {
			log.Println("WEBSOCKET. INCOMING MESSAGE unmarhsal", err)
			continue
		}
		chm.User = c.u
		ch <- &chm
	}
}

type EventMessage struct {
	Type  string       `json:"type"`
	Event *types.Event `json:"event"`
}

func (wr *websocketRoom) sendMessageToChat(id int64, msg *ChatMessage) {
	wr.mux.RLock()
	defer wr.mux.RUnlock()

	if c, ok := wr.Clients[id]; !ok {
		return
	} else {
		if err := c.conn.WriteJSON(msg); err != nil {
			log.Println("WEBSOCKET. WRITE MESSAGE TO CHAT", err)
		}
	}
}

func (wr *websocketRoom) sendEventInvitationToClient(id int64, e *types.Event) {
	wr.mux.RLock()
	defer wr.mux.RUnlock()

	if c, ok := wr.Clients[id]; !ok {
		return
	} else {
		if err := c.conn.WriteJSON(EventMessage{"event", e}); err != nil {
			log.Println("WEBSOCKET. WRITE MESSAGE", err)
		}
	}
}
