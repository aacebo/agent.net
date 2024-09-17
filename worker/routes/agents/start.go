package agents

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aacebo/agent.net/amqp"
	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/rabbitmq/amqp091-go"
)

func Start(ctx context.Context) func(amqp091.Delivery) {
	log := logger.New("agent.net/worker/agents/start")
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

		var parent *models.Agent

		if agent.ParentID != nil {
			p, exists := agents.GetByID(*agent.ParentID)

			if exists {
				parent = &p
			}
		}

		if parent != nil && parent.Address == nil {
			log.Warn("agent created but not started due to parent not being started")
			m.Nack(false, true)
			return
		}

		settings, err := json.Marshal(agent.Settings)

		if err != nil {
			log.Error(err.Error())
			m.Nack(false, true)
			return
		}

		env := []string{
			fmt.Sprintf("AGENT_ID=%s", agent.ID),
			fmt.Sprintf("AGENT_CLIENT_ID=%s", agent.ClientID),
			fmt.Sprintf("AGENT_CLIENT_SECRET=%s", agent.ClientSecret),
			fmt.Sprintf("AGENT_NAME=%s", agent.Name),
			fmt.Sprintf("AGENT_DESCRIPTION=%s", agent.Description),
			fmt.Sprintf("AGENT_INSTRUCTIONS=%s", *agent.Instructions),
			fmt.Sprintf("AGENT_SETTINGS=%s", string(settings)),
		}

		if parent != nil {
			env = append(
				env,
				fmt.Sprintf("AGENT_ADDRESS=%s", *parent.Address),
			)
		} else {
			env = append(env, fmt.Sprintf("AGENT_ADDRESS=%s", "agent-net.ngrok.io")) // if no parent, agent should connect to main server
		}

		if agent.ContainerID == nil {
			res, err := docker.ContainerCreate(context.Background(), &container.Config{
				Image:        "agent.net/agent",
				ExposedPorts: nat.PortSet{"8080": struct{}{}},
				Env:          env,
			}, nil, nil, nil, fmt.Sprintf("agent.net-agent-%s", agent.ID))

			if err != nil {
				log.Error(err.Error())
				m.Nack(false, true)
				return
			}

			agent.ContainerID = &res.ID
			agent = agents.Update(agent)
		}

		if err := docker.ContainerStart(context.Background(), *agent.ContainerID, container.StartOptions{}); err != nil {
			log.Error(err.Error())
			m.Nack(false, true)
			return
		}

		log.Debug(fmt.Sprintf("agent started: %s", agent.ID))
		m.Ack(false)
	}
}
