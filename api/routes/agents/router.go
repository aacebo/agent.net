package agents

import (
	"github.com/aacebo/agent.net/api/common"
	"github.com/aacebo/agent.net/api/middleware"
	"github.com/go-chi/chi/v5"
)

func New(r chi.Router, ctx common.Context) {
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
}
