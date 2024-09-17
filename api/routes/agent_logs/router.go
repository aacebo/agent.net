package agentLogs

import (
	"context"

	"github.com/aacebo/agent.net/api/middleware"
	"github.com/go-chi/chi/v5"
)

func New(r chi.Router, ctx context.Context) {
	r.With(
		middleware.WithAgent(ctx),
	).Get(
		"/agents/{agent_id}/logs",
		GetByAgent(ctx),
	)

	r.With(
		middleware.WithAgent(ctx),
		middleware.WithBody[CreateBody](ctx, "/agent_logs/create"),
	).Post(
		"/agents/{agent_id}/logs",
		Create(ctx),
	)
}
