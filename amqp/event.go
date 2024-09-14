package amqp

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Event[T any] struct {
	ID        string    `json:"id"`
	Body      T         `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func NewEvent[T any](body T) Event[T] {
	return Event[T]{
		ID:        uuid.NewString(),
		Body:      body,
		CreatedAt: time.Now(),
	}
}

func (self Event[T]) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self *Event[T]) Decode(data []byte) error {
	dec := gob.NewDecoder(bytes.NewBuffer(data))
	return dec.Decode(self)
}
