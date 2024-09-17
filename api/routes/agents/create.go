package agents

import (
	"context"
	"net/http"
	"strings"

	"github.com/aacebo/agent.net/amqp"
	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/aacebo/agent.net/core/utils"
	"github.com/go-chi/render"
)

type CreateBody struct {
	Name         *string              `json:"name,omitempty"`
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

		agent.Name = strings.ToLower(utils.RandString(10))

		if body.Name != nil {
			agent.Name = *body.Name
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
