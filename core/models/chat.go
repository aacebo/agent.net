package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID        string    `json:"id"`
	AgentID   string    `json:"agent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewChat() Chat {
	return Chat{
		ID: uuid.NewString(),
	}
}

func (self Chat) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}
