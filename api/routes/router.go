package routes

import (
	"context"

	agentLogs "github.com/aacebo/agent.net/api/routes/agent_logs"
	"github.com/aacebo/agent.net/api/routes/agents"
	"github.com/go-chi/chi/v5"
)

func New(ctx context.Context) *chi.Mux {
	r := chi.NewRouter()

	agents.New(r, ctx)
	agentLogs.New(r, ctx)
	r.HandleFunc("/sockets", Socket(ctx))

	return r
}
