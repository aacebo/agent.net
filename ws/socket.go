package ws

import (
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/aacebo/agent.net/core/logger"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Socket struct {
	ID        string
	CreatedAt time.Time

	conn *websocket.Conn
	log  *slog.Logger
	mu   sync.RWMutex
}

func newSocket(conn *websocket.Conn) *Socket {
	id := uuid.NewString()
	socket := Socket{
		ID:        id,
		CreatedAt: time.Now(),

		conn: conn,
		log:  logger.New(fmt.Sprintf("agent.net/socket/%s", id)),
		mu:   sync.RWMutex{},
	}

	go socket.onPing()
	return &socket
}

func (self *Socket) Read() (Message, error) {
	self.mu.RLock()
	defer self.mu.RUnlock()

	msg := Message{}
	err := self.conn.ReadJSON(&msg)

	if err != nil {
		self.log.Warn(err.Error())
	}

	return msg, err
}

func (self *Socket) Send(msg Message) error {
	self.mu.Lock()
	defer self.mu.Unlock()

	err := self.conn.WriteJSON(msg)

	if err != nil {
		self.log.Warn(err.Error())
		self.conn.Close()
	}

	return err
}

func (self *Socket) onPing() {
	for range time.Tick(20 * time.Second) {
		err := self.conn.WriteMessage(websocket.PingMessage, []byte{})

		if err != nil {
			self.log.Warn(err.Error())
			self.conn.Close()
			return
		}
	}
}
