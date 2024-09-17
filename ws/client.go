package ws

import (
	"log/slog"
	"net/http"
	"sync"

	"github.com/aacebo/agent.net/core/logger"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn  *websocket.Conn
	log   *slog.Logger
	read  sync.Mutex
	write sync.Mutex
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

	self.conn = conn
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
		return msg, err
	}

	self.log.Debug(msg.String())
	return msg, err
}

func (self *Client) Send(msg Message) error {
	self.write.Lock()
	defer self.write.Unlock()

	err := self.conn.WriteJSON(msg)

	if err != nil {
		self.conn.Close()
		return err
	}

	self.log.Debug(msg.String())
	return err
}
