package agents

import (
	"context"

	"github.com/aacebo/agent.net/api/middleware"
	"github.com/go-chi/chi/v5"
)

func New(r chi.Router, ctx context.Context) {
	r.HandleFunc(
		"/sockets",
		Handler(ctx),
	)

	r.With(
		middleware.WithBody[CreateBody](ctx, "/agents/create"),
	).Post(
		"/agents",
		Create(ctx),
	)

	r.With(
		middleware.WithAgent(ctx),
		middleware.WithBody[CreateBody](ctx, "/agents/create"),
	).Post(
		"/agents/{agent_id}/agents",
		Create(ctx),
	)
}
