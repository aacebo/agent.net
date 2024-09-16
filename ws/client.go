package ws

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	mu   sync.RWMutex
}

func NewClient() *Client {
	return &Client{
		mu: sync.RWMutex{},
	}
}

func (self *Client) Connect(url string, headers http.Header) error {
	self.mu.Lock()
	defer self.mu.Unlock()

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
	self.mu.RLock()
	defer self.mu.RUnlock()

	msg := Message{}
	err := self.conn.ReadJSON(&msg)

	if err != nil {
		return msg, err
	}

	fmt.Println(msg)
	return msg, err
}

func (self *Client) Send(msg Message) error {
	self.mu.Lock()
	defer self.mu.Unlock()

	err := self.conn.WriteJSON(msg)

	if err != nil {
		self.conn.Close()
		return err
	}

	return err
}
