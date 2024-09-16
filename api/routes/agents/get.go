package agents

import (
	"context"
	"net/http"

	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/go-chi/render"
)

func Get(ctx context.Context) http.HandlerFunc {
	agents := ctx.Value("repos.agents").(repos.IAgentsRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		parent, isChild := r.Context().Value("agent").(models.Agent)
		arr := []models.Agent{}

		if isChild {
			arr = agents.GetEdges(parent.ID)
		} else {
			arr = agents.Get()
		}

		render.JSON(w, r, arr)
	}
}
