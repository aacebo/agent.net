package routes

import (
	"github.com/aacebo/agent.net/api/common"
	"github.com/aacebo/agent.net/api/routes/agents"
	"github.com/go-chi/chi/v5"
)

func New(ctx common.Context) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		agents.New(r, ctx)
	})

	return r
}
