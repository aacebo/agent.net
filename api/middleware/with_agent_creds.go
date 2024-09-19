package middleware

import (
	"context"
	"net/http"

	"github.com/aacebo/agent.net/core/repos"
	"github.com/go-chi/render"
)

func WithAgentCreds(ctx context.Context) func(http.Handler) http.Handler {
	agents := ctx.Value("repos.agents").(repos.IAgentsRepository)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			clientId := r.Header.Get("X_CLIENT_ID")
			clientSecret := r.Header.Get("X_CLIENT_SECRET")
			agent, exists := agents.GetByCreds(clientId, clientSecret)

			if !exists {
				render.Status(r, http.StatusUnauthorized)
				render.JSON(w, r, "unauthorized")
				return
			}

			ctx = context.WithValue(ctx, "agent", agent)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
