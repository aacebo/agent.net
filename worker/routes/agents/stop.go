package agents

import (
	"context"
	"fmt"

	"github.com/aacebo/agent.net/amqp"
	"github.com/aacebo/agent.net/containers"
	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
	"github.com/rabbitmq/amqp091-go"
)

func Stop(ctx context.Context) func(amqp091.Delivery) {
	log := logger.New("worker/agents/stop")
	client := ctx.Value("containers").(containers.Client)

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

		if agent.Status == models.AGENT_STATUS_DOWN {
			log.Warn("agent is already stopped")
			m.Nack(false, false)
			return
		}

		if err := client.Stop(*agent.ContainerID); err != nil {
			log.Error(err.Error())
			m.Nack(false, true)
			return
		}

		log.Debug(fmt.Sprintf("agent stopped: %s", agent.ID))
		m.Ack(false)
	}
}
