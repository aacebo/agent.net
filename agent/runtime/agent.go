package runtime

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
	"sync"
	"time"

	"github.com/aacebo/agent.net/agent/client"
	"github.com/aacebo/agent.net/core/models"
	_slices "github.com/aacebo/agent.net/core/utils/slices"
	"github.com/aacebo/agent.net/ws"
	"github.com/gorilla/websocket"
)

type Agent struct {
	ID            string
	Address       string
	ParentAddress string
	Name          string
	Description   string
	ClientID      string
	ClientSecret  string
	StartedAt     time.Time
	Edges         []models.AgentStat

	log      *client.LoggerClient
	parent   *ws.Client
	children *ws.Sockets
	mu       sync.RWMutex
}

func NewAgent(
	id string,
	address string,
	parentAddress string,
	name string,
	description string,
	clientId string,
	clientSecret string,
	startedAt time.Time,
) *Agent {
	return &Agent{
		ID:            id,
		Address:       address,
		ParentAddress: parentAddress,
		Description:   description,
		ClientID:      clientId,
		ClientSecret:  clientSecret,
		StartedAt:     startedAt,
		Edges:         []models.AgentStat{},

		log:      client.NewLogger(fmt.Sprintf("agent.net/%s/runtime", name)),
		parent:   ws.NewClient(),
		children: ws.NewSockets(),
		mu:       sync.RWMutex{},
	}
}

func (self *Agent) Logger() *client.LoggerClient {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.log
}

func (self *Agent) Add(r *http.Request, conn *websocket.Conn) *ws.Socket {
	self.mu.Lock()
	defer self.mu.Unlock()

	id := r.Header.Get("X_AGENT_ID")
	socket := self.children.Add(conn)
	agent := models.AgentStat{
		ID:          r.Header.Get("X_AGENT_ID"),
		SocketID:    socket.ID,
		Name:        r.Header.Get("X_AGENT_NAME"),
		Description: r.Header.Get("X_AGENT_DESCRIPTION"),
		StartedAt:   time.Now(),
	}

	i := slices.IndexFunc(self.Edges, func(a models.AgentStat) bool {
		return id == a.ID
	})

	if i == -1 {
		self.Edges = append(self.Edges, agent)
		return socket
	}

	self.Edges[i] = agent
	return socket
}

func (self *Agent) Remove(socketId string) error {
	self.mu.Lock()
	defer self.mu.Unlock()

	socket := self.children.GetByID(socketId)

	if socket == nil {
		return nil
	}

	if err := socket.Close(); err != nil {
		return err
	}

	self.children.Del(socketId)
	return nil
}

func (self *Agent) SendToParent(msg ws.Message) error {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.parent.Send(msg)
}

func (self *Agent) SendToAgent(agentId string, msg ws.Message) error {
	self.mu.RLock()
	defer self.mu.RUnlock()

	agent, exists := _slices.Find(self.Edges, func(child models.AgentStat) bool {
		return child.ID == agentId
	})

	if !exists {
		return errors.New("agent not found")
	}

	socket := self.children.GetByID(agent.SocketID)

	if socket == nil {
		return errors.New("agent socket not found")
	}

	socket.Send(msg)
	return nil
}

func (self *Agent) SendToAgents(msg ws.Message) error {
	self.mu.RLock()
	defer self.mu.RUnlock()

	for _, agent := range self.Edges {
		socket := self.children.GetByID(agent.SocketID)

		if socket == nil {
			return errors.New("agent socket not found")
		}

		socket.Send(msg)
	}

	return nil
}
