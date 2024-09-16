package agents

import (
	"context"
	"net/http"

	"github.com/aacebo/agent.net/amqp"
	"github.com/aacebo/agent.net/core/models"
	"github.com/go-chi/render"
)

func Stop(ctx context.Context) http.HandlerFunc {
	amqp := ctx.Value("amqp").(*amqp.Client)

	return func(w http.ResponseWriter, r *http.Request) {
		agent := r.Context().Value("agent").(models.Agent)

		if agent.Status == models.AGENT_STATUS_DOWN {
			render.Status(r, http.StatusConflict)
			render.JSON(w, r, "agent already stopped")
			return
		}

		amqp.Publish("agents", "stop", agent)
		render.JSON(w, r, map[string]any{
			"ok": true,
		})
	}
}
