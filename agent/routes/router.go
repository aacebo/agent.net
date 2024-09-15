package routes

import (
	"context"

	"github.com/go-chi/chi/v5"
)

func New(ctx context.Context) *chi.Mux {
	r := chi.NewRouter()

	r.HandleFunc("/sockets", Socket(ctx))

	return r
}
