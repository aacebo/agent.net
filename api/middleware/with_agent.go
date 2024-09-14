package middleware

import (
	"context"
	"net/http"

	"github.com/aacebo/agent.net/core/repos"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func WithAgent(ctx context.Context) func(http.Handler) http.Handler {
	agents := ctx.Value("repos.agents").(repos.IAgentsRepository)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			id := chi.URLParam(r, "agent_id")

			if _, err := uuid.Parse(id); err != nil {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, "invalid `agent_id`")
				return
			}

			agent, exists := agents.GetByID(id)

			if !exists {
				render.Status(r, http.StatusNotFound)
				render.JSON(w, r, "agent not found")
				return
			}

			ctx = context.WithValue(ctx, "agent", agent)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
