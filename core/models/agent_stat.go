package models

import (
	"encoding/json"
	"time"
)

type AgentStat struct {
	ID          string    `json:"id"`
	SocketID    string    `json:"socket_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IPAddress   string    `json:"ip_address"`
	StartedAt   time.Time `json:"started_at"`
}

func (self AgentStat) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}
