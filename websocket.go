package potato

import (
	"log"
	"github.com/gorilla/websocket"
	"sync"
	"github.com/flameous/PotatoPartyBackend/types"
	"net/http"
	"time"
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
}

func newWebsocketRoom(out chan []int64) *websocketRoom {
	wr := websocketRoom{make(map[int64]*Client), &sync.RWMutex{}}
	go wr.cycle(out)
	return &wr
}

func (wr *websocketRoom) cycle(out chan []int64) {
	for {
		time.Sleep(time.Second * 30)
		clients := wr.readAllClients()
		out <- clients
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
	wr.Clients[u.ID] = &Client{u, conn}
}

type Client struct {
	u    *types.User
	conn *websocket.Conn
}

type MessageEvent struct {
	Type  string       `json:"type"`
	Event *types.Event `json:"event"`
}

func (wr *websocketRoom) sendEventInvitationToClient(id int64, e *types.Event) {
	wr.mux.RLock()
	defer wr.mux.RUnlock()

	if c, ok := wr.Clients[id]; !ok {
		return
	} else {
		if err := c.conn.WriteJSON(MessageEvent{"event", e}); err != nil {
			log.Println("WEBSOCKET. WRITE MESSAGE", err)
		}
	}
}
