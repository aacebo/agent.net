package models

import (
	"encoding/json"
	"time"
)

type Stat struct {
	ID          string      `json:"id"`
	Description string      `json:"description"`
	StartedAt   time.Time   `json:"started_at"`
	Edges       []AgentStat `json:"edges"`
}

func (self Stat) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}
