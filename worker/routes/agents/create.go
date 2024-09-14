package agents

import (
	"context"
	"fmt"

	"github.com/aacebo/agent.net/amqp"
	"github.com/aacebo/agent.net/core/models"
	"github.com/rabbitmq/amqp091-go"
)

func Create(ctx context.Context) func(amqp091.Delivery) {
	return func(m amqp091.Delivery) {
		event := amqp.Event[models.Agent]{}

		if err := event.Decode(m.Body); err != nil {
			m.Nack(false, true)
			return
		}

		fmt.Println(event)
		m.Ack(false)
	}
}
