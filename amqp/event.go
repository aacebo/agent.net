package amqp

import "time"

type Event struct {
	ID        string    `json:"id"`
	Body      any       `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
