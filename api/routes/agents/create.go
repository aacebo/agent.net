package agents

import (
	"net/http"

	"github.com/aacebo/agent.net/api/common"
	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/go-chi/render"
)

type CreateBody struct {
	Description  string               `json:"description"`
	Instructions *string              `json:"instructions,omitempty"`
	Settings     models.AgentSettings `json:"settings"`
}

func Create(ctx common.Context) http.HandlerFunc {
	agents := ctx.Value("repos.agents").(repos.IAgentRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Context().Value("body").(CreateBody)

		agent := models.NewAgent()
		agent.Description = body.Description
		agent.Instructions = body.Instructions
		agent.Settings = body.Settings
		agent = agents.Create(agent)

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, agent)
	}
}
