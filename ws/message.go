package ws

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID     string      `json:"id"`
	Type   MessageType `json:"type"`
	Body   any         `json:"body,omitempty"`
	SentAt time.Time   `json:"sent_at"`
}

func NewMessage(t MessageType, body any) Message {
	return Message{
		ID:     uuid.NewString(),
		Type:   t,
		Body:   body,
		SentAt: time.Now(),
	}
}

func (self Message) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}
