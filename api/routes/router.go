package routes

import (
	"context"

	"github.com/aacebo/agent.net/api/middleware"
	agentLogs "github.com/aacebo/agent.net/api/routes/agent_logs"
	"github.com/aacebo/agent.net/api/routes/agents"
	"github.com/aacebo/agent.net/api/routes/messages"
	"github.com/go-chi/chi/v5"
)

func New(ctx context.Context) *chi.Mux {
	r := chi.NewRouter()

	agents.New(r, ctx)
	agentLogs.New(r, ctx)
	messages.New(r, ctx)

	r.With(
		middleware.WithAgentCreds(ctx),
	).HandleFunc(
		"/sockets",
		Socket(ctx),
	)

	return r
}
