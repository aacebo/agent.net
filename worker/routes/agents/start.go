package agents

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aacebo/agent.net/amqp"
	"github.com/aacebo/agent.net/containers"
	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/rabbitmq/amqp091-go"
)

func Start(ctx context.Context) func(amqp091.Delivery) {
	log := logger.New("agent.net/worker/agents/start")
	client := ctx.Value("containers").(containers.Client)
	agents := ctx.Value("repos.agents").(repos.IAgentsRepository)
	port := 8080

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
			id, err := client.Create(containers.ContainerCreateArgs{
				Image:     "agent.net/agent",
				Name:      fmt.Sprintf("agent.net-agent-%s", agent.ID),
				Port:      port,
				IPAddress: "0.0.0.0",
				Env:       env,
			})

			if err != nil {
				log.Error(err.Error())
				m.Nack(false, false)
				return
			}

			agent.ContainerID = &id
			agent = agents.Update(agent)
		}

		if err := client.Start(*agent.ContainerID); err != nil {
			log.Error(err.Error())
			m.Nack(false, true)
			return
		}

		port++
		log.Debug(fmt.Sprintf("agent started: %s", agent.ID))
		m.Ack(false)
	}
}
