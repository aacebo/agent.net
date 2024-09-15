package routes

import (
	"context"

	"github.com/aacebo/agent.net/api/routes/agents"
	"github.com/go-chi/chi/v5"
)

func New(ctx context.Context) *chi.Mux {
	r := chi.NewRouter()

	agents.New(r, ctx)
	r.HandleFunc("/sockets", Socket(ctx))

	return r
}
