package ws

import (
	"math/rand"
	"slices"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Sockets struct {
	mu      sync.RWMutex
	sockets []*Socket
}

func NewSockets() *Sockets {
	return &Sockets{
		mu:      sync.RWMutex{},
		sockets: []*Socket{},
	}
}

func (self *Sockets) Get() *Socket {
	self.mu.RLock()
	defer self.mu.RUnlock()

	seed := rand.New(rand.NewSource(time.Now().Unix()))
	return self.sockets[seed.Intn(len(self.sockets))]
}

func (self *Sockets) GetByID(id string) *Socket {
	self.mu.RLock()
	defer self.mu.RUnlock()

	i := slices.IndexFunc(self.sockets, func(s *Socket) bool {
		return s.ID == id
	})

	if i == -1 {
		return nil
	}

	return self.sockets[i]
}

func (self *Sockets) Add(conn *websocket.Conn) *Socket {
	self.mu.Lock()
	defer self.mu.Unlock()

	socket := newSocket(conn)
	self.sockets = append(self.sockets, socket)
	return socket
}

func (self *Sockets) Del(id string) {
	self.mu.Lock()
	defer self.mu.Unlock()

	self.sockets = slices.DeleteFunc(self.sockets, func(s *Socket) bool {
		return s.ID == id
	})
}

func (self *Sockets) Send(msg Message) error {
	self.mu.Lock()
	defer self.mu.Unlock()

	for _, socket := range self.sockets {
		if err := socket.Send(msg); err != nil {
			return err
		}
	}

	return nil
}
