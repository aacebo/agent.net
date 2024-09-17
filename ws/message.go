package ws

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID     string      `json:"id"`
	ToID   string      `json:"to_id"`
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

func NewTextMessage(text string) Message {
	return Message{
		ID:     uuid.NewString(),
		Type:   TEXT_MESSAGE_TYPE,
		Body:   text,
		SentAt: time.Now(),
	}
}

func NewConnectedMessage(body ConnectedMessageBody) Message {
	return Message{
		ID:     uuid.NewString(),
		Type:   CONNECTED_MESSAGE_TYPE,
		Body:   body,
		SentAt: time.Now(),
	}
}

func NewDisconnectedMessage(id string) Message {
	return Message{
		ID:     uuid.NewString(),
		Type:   DISCONNECTED_MESSAGE_TYPE,
		Body:   id,
		SentAt: time.Now(),
	}
}

func (self Message) WithToID(toId string) Message {
	self.ToID = toId
	return self
}

func (self Message) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

type ConnectedMessageBody map[string]any

func (self ConnectedMessageBody) ID() string {
	return self["id"].(string)
}

func (self ConnectedMessageBody) Address() string {
	return self["address"].(string)
}
