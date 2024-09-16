package client

import (
	"log/slog"
	"net/http"
	"slices"
	"sync"
	"time"

	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/ws"
)

type Client struct {
	url     string
	stat    models.Stat
	client  *ws.Client
	sockets *ws.Sockets
	log     *slog.Logger
	mu      sync.RWMutex
}

func New(
	id string,
	url string,
	description string,
	startedAt time.Time,
	sockets *ws.Sockets,
) *Client {
	return &Client{
		url: url,
		stat: models.Stat{
			ID:          id,
			Description: description,
			StartedAt:   startedAt,
			Edges:       []models.AgentStat{},
		},
		client:  ws.NewClient(),
		sockets: sockets,
		log:     logger.New("agent.net/agent"),
		mu:      sync.RWMutex{},
	}
}

func (self *Client) Stat() models.Stat {
	return self.stat
}

func (self *Client) Listen(clientId string, clientSecret string) {
	header := http.Header{}
	header.Add("client_id", clientId)
	header.Add("client_secret", clientSecret)

	if err := self.client.Connect(self.url, header); err != nil {
		panic(err)
	}

	defer self.client.Close()
	self.log.Info("connected...")

	for {
		message, err := self.client.Read()

		if err != nil || !message.Type.Valid() {
			self.log.Warn(err.Error())
			return
		}

		switch message.Type {
		case ws.QUERY_MESSAGE_TYPE:
			body := message.Body.(ws.QueryMessageBody)

			switch body.Name {
			case "stat":
				self.client.Send(ws.NewQueryResponseMessage("stat", self.stat))

				for _, agent := range self.stat.Edges {
					socket := self.sockets.GetByID(agent.SocketID)

					if socket == nil {
						continue
					}

					socket.Send(ws.NewQueryMessage("stat"))
				}
			}
		}
	}
}

func (self *Client) Send(message ws.Message) error {
	self.mu.Lock()
	defer self.mu.Unlock()
	return self.client.Send(message)
}

func (self *Client) SetAgent(agent models.AgentStat) {
	self.mu.Lock()
	defer self.mu.Unlock()

	i := slices.IndexFunc(self.stat.Edges, func(a models.AgentStat) bool {
		return agent.ID == a.ID
	})

	if i == -1 {
		self.stat.Edges = append(self.stat.Edges, agent)
		return
	}

	self.stat.Edges[i] = agent
}
