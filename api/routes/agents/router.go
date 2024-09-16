package agents

import (
	"context"

	"github.com/aacebo/agent.net/api/middleware"
	"github.com/go-chi/chi/v5"
)

func New(r chi.Router, ctx context.Context) {
	r.Get("/agents", Get(ctx))

	r.With(
		middleware.WithAgent(ctx),
	).Get(
		"/agents/{agent_id}/agents",
		Get(ctx),
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

	r.With(
		middleware.WithAgent(ctx),
	).Post(
		"/agents/{agent_id}/start",
		Start(ctx),
	)

	r.With(
		middleware.WithAgent(ctx),
	).Post(
		"/agents/{agent_id}/stop",
		Stop(ctx),
	)
}
