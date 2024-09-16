package ws

import (
	"encoding/json"
	"errors"
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

func NewTextMessage(text string) Message {
	return Message{
		ID:     uuid.NewString(),
		Type:   TEXT_MESSAGE_TYPE,
		Body:   TextMessageBody{text},
		SentAt: time.Now(),
	}
}

func NewQueryMessage(name string) Message {
	return Message{
		ID:     uuid.NewString(),
		Type:   QUERY_MESSAGE_TYPE,
		Body:   QueryMessageBody{name},
		SentAt: time.Now(),
	}
}

func NewQueryResponseMessage(name string, body any) Message {
	return Message{
		ID:   uuid.NewString(),
		Type: QUERY_RESPONSE_MESSAGE_TYPE,
		Body: QueryResponseMessageBody{
			Name: name,
			Body: body,
		},
		SentAt: time.Now(),
	}
}

func (self Message) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(self)
}

func (self *Message) UnmarshalJSON(data []byte) error {
	d := map[string]any{}
	err := json.Unmarshal(data, &d)

	if err != nil {
		return err
	}

	sentAt, err := time.Parse(time.RFC3339, d["sent_at"].(string))

	if err != nil {
		return err
	}

	self.ID = d["id"].(string)
	self.Type = MessageType(d["type"].(string))
	self.SentAt = sentAt

	if !self.Type.Valid() {
		return errors.New("invalid message payload")
	}

	bodyData, err := json.Marshal(d["body"].(map[string]any))

	if err != nil {
		return nil
	}

	switch self.Type {
	case TEXT_MESSAGE_TYPE:
		body := TextMessageBody{}
		err = json.Unmarshal(bodyData, &body)

		if err != nil {
			return err
		}

		self.Body = body
	case QUERY_MESSAGE_TYPE:
		body := QueryMessageBody{}
		err = json.Unmarshal(bodyData, &body)

		if err != nil {
			return err
		}

		self.Body = body
	case QUERY_RESPONSE_MESSAGE_TYPE:
		body := QueryResponseMessageBody{}
		err = json.Unmarshal(bodyData, &body)

		if err != nil {
			return err
		}

		self.Body = body
	}

	return nil
}

type TextMessageBody struct {
	Text string `json:"text"`
}

func (self TextMessageBody) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

type QueryMessageBody struct {
	Name string `json:"name"`
}

func (self QueryMessageBody) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

type QueryResponseMessageBody struct {
	Name string `json:"name"`
	Body any    `json:"body"`
}

func (self QueryResponseMessageBody) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}
