package amqp

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"time"
)

type Event struct {
	ID        string    `json:"id"`
	Body      any       `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func (self Event) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self *Event) Decode(data []byte) error {
	dec := gob.NewDecoder(bytes.NewBuffer(data))
	return dec.Decode(self)
}
