package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        string    `json:"id"`
	ChatID    string    `json:"chat_id"`
	FromID    string    `json:"from_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewMessage() Message {
	return Message{
		ID: uuid.NewString(),
	}
}

func (self Message) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}
