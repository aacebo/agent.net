package models

import (
	"encoding/json"
	"time"

	"github.com/aacebo/agent.net/core/utils"
	"github.com/google/uuid"
)

type Agent struct {
	ID           string        `json:"id"`
	ParentID     *string       `json:"parent_id,omitempty"`
	ContainerID  *string       `json:"container_id,omitempty"`
	Description  string        `json:"description"`
	Status       AgentStatus   `json:"status"`
	Instructions *string       `json:"instructions,omitempty"`
	URL          *string       `json:"url,omitempty"`
	ClientID     string        `json:"-"`
	ClientSecret Secret        `json:"-"`
	Settings     AgentSettings `json:"settings"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

func NewAgent() Agent {
	return Agent{
		ID:           uuid.NewString(),
		Status:       AGENT_STATUS_DOWN,
		ClientID:     utils.RandString(20),
		ClientSecret: Secret(utils.RandString(32)),
		Settings:     AgentSettings{},
	}
}

func (self Agent) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}
