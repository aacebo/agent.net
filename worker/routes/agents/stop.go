package agents

import (
	"context"

	"github.com/aacebo/agent.net/amqp"
	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/rabbitmq/amqp091-go"
)

func Stop(ctx context.Context) func(amqp091.Delivery) {
	log := logger.New("agent.net/worker/agents/deploy")
	docker := ctx.Value("docker").(*client.Client)
	agents := ctx.Value("repos.agents").(repos.IAgentsRepository)

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

		err := docker.ContainerStop(
			context.Background(),
			*agent.ContainerID,
			container.StopOptions{},
		)

		if err != nil {
			log.Error(err.Error())
			m.Nack(false, true)
			return
		}

		agent.Status = models.AGENT_STATUS_DOWN
		agent = agents.Update(agent)

		m.Ack(false)
	}
}
