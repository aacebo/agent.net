package agents

import (
	"context"
	"net/http"

	"github.com/aacebo/agent.net/amqp"
	"github.com/aacebo/agent.net/core/models"
	"github.com/go-chi/render"
)

func Start(ctx context.Context) http.HandlerFunc {
	amqp := ctx.Value("amqp").(*amqp.Client)

	return func(w http.ResponseWriter, r *http.Request) {
		agent := r.Context().Value("agent").(models.Agent)

		if agent.Status == models.AGENT_STATUS_UP {
			render.Status(r, http.StatusConflict)
			render.JSON(w, r, "agent already started")
			return
		}

		amqp.Publish("agents", "start", agent)
		render.JSON(w, r, map[string]any{
			"ok": true,
		})
	}
}
