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

	conn  *websocket.Conn
	log   *slog.Logger
	read  sync.Mutex
	write sync.Mutex
}

func newSocket(conn *websocket.Conn) *Socket {
	id := uuid.NewString()
	socket := Socket{
		ID:        id,
		CreatedAt: time.Now(),

		conn:  conn,
		log:   logger.New(fmt.Sprintf("agent.net/socket/%s", id)),
		read:  sync.Mutex{},
		write: sync.Mutex{},
	}

	go socket.ping()
	return &socket
}

func (self *Socket) IPAddress() string {
	return self.conn.RemoteAddr().String()
}

func (self *Socket) Read() (Message, error) {
	self.read.Lock()
	defer self.read.Unlock()

	msg := Message{}
	err := self.conn.ReadJSON(&msg)

	if err != nil {
		self.log.Warn(err.Error())
		return msg, err
	}

	self.log.Debug(msg.String())
	return msg, err
}

func (self *Socket) Send(msg Message) error {
	self.write.Lock()
	defer self.write.Unlock()

	err := self.conn.WriteJSON(msg)

	if err != nil {
		self.log.Warn(err.Error())
		self.conn.Close()
		return err
	}

	self.log.Debug(msg.String())
	return err
}

func (self *Socket) ping() {
	for range time.Tick(20 * time.Second) {
		self.write.Lock()

		err := self.conn.WriteMessage(websocket.PingMessage, []byte{})

		if err != nil {
			self.log.Warn(err.Error())
			self.conn.Close()
			self.write.Unlock()
			return
		}

		self.write.Unlock()
	}
}
