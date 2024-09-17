package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type AgentLog struct {
	ID        string    `json:"id"`
	AgentID   string    `json:"agent_id"`
	Level     LogLevel  `json:"level"`
	Text      string    `json:"text"`
	Data      Map[any]  `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAgentLog() AgentLog {
	return AgentLog{
		ID: uuid.NewString(),
	}
}

func (self AgentLog) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}
