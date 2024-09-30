package agents

import (
	"context"
	"fmt"

	"github.com/aacebo/agent.net/amqp"
	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
	"github.com/rabbitmq/amqp091-go"
)

func Create(ctx context.Context) func(amqp091.Delivery) {
	log := logger.New("worker/agents/create")

	return func(m amqp091.Delivery) {
		event := amqp.Event{}

		if err := event.Decode(m.Body); err != nil {
			log.Error(err.Error())
			m.Nack(false, true)
			return
		}

		agent, ok := event.Body.(models.Agent)

		if !ok {
			log.Error("failed to decode event body: %v", event.Body)
			m.Nack(false, true)
			return
		}

		log.Debug(fmt.Sprintf("agent created: %s", agent.ID))
		m.Ack(false)
	}
}
