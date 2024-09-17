package ws

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/aacebo/agent.net/core/logger"
	"github.com/gorilla/websocket"
)

type Client struct {
	url               string
	headers           http.Header
	conn              *websocket.Conn
	log               *slog.Logger
	reconnectAttempts int
	read              sync.Mutex
	write             sync.Mutex
}

func NewClient() *Client {
	return &Client{
		log:   logger.New("agent.net/agent/socket"),
		read:  sync.Mutex{},
		write: sync.Mutex{},
	}
}

func (self *Client) Connect(url string, headers http.Header) error {
	self.write.Lock()
	defer self.write.Unlock()

	conn, _, err := websocket.DefaultDialer.Dial(url, headers)

	if err != nil {
		return err
	}

	self.url = url
	self.headers = headers
	self.conn = conn
	self.reconnectAttempts = 0
	return nil
}

func (self *Client) Close() error {
	return self.conn.Close()
}

func (self *Client) Read() (Message, error) {
	self.read.Lock()
	defer self.read.Unlock()

	msg := Message{}
	err := self.conn.ReadJSON(&msg)

	if err != nil {
		self.reconnect()
		return self.Read()
	}

	return msg, err
}

func (self *Client) Send(msg Message) error {
	self.write.Lock()
	defer self.write.Unlock()

	err := self.conn.WriteJSON(msg)

	if err != nil {
		self.reconnect()
		return nil
	}

	return err
}

func (self *Client) reconnect() {
	self.Close()

	for {
		self.reconnectAttempts++
		ms := (500 * self.reconnectAttempts)
		time.Sleep(time.Duration(ms) * time.Millisecond)
		self.log.Debug(fmt.Sprintf("attempting reconnect %d...", self.reconnectAttempts))

		if err := self.Connect(self.url, self.headers); err == nil {
			break
		}
	}
}
