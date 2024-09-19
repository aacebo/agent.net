package agents

import (
	"context"
	"net/http"

	"github.com/aacebo/agent.net/amqp"
	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/go-chi/render"
)

type UpdateBody struct {
	Description  *string               `json:"description,omitempty"`
	Instructions *string               `json:"instructions,omitempty"`
	Settings     *models.AgentSettings `json:"settings,omitempty"`
	Position     *models.Position      `json:"position,omitempty"`
}

func Update(ctx context.Context) http.HandlerFunc {
	amqp := ctx.Value("amqp").(*amqp.Client)
	agents := ctx.Value("repos.agents").(repos.IAgentsRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		agent := r.Context().Value("agent").(models.Agent)
		body := r.Context().Value("body").(UpdateBody)

		if body.Description != nil {
			agent.Description = *body.Description
		}

		if body.Instructions != nil {
			agent.Instructions = body.Instructions
		}

		if body.Settings != nil {
			agent.Settings = *body.Settings
		}

		if body.Position != nil {
			agent.Position = *body.Position
		}

		agent = agents.Update(agent)
		amqp.Publish("agents", "update", agent)
		render.JSON(w, r, agent)
	}
}
