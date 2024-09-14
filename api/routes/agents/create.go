package agents

import (
	"context"
	"net/http"

	"github.com/aacebo/agent.net/amqp"
	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/go-chi/render"
)

type CreateBody struct {
	Description  string               `json:"description"`
	Instructions *string              `json:"instructions,omitempty"`
	Settings     models.AgentSettings `json:"settings"`
}

func Create(ctx context.Context) http.HandlerFunc {
	amqp := ctx.Value("amqp").(*amqp.Client)
	agents := ctx.Value("repos.agents").(repos.IAgentsRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		parent, isChild := r.Context().Value("agent").(models.Agent)
		body := r.Context().Value("body").(CreateBody)

		agent := models.NewAgent()

		if isChild {
			agent.ParentID = &parent.ID
		}

		agent.Description = body.Description
		agent.Instructions = body.Instructions
		agent.Settings = body.Settings
		agent = agents.Create(agent)
		amqp.Publish("agents", "create", agent)

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, agent)
	}
}
