package messages

import (
	"context"

	"github.com/aacebo/agent.net/api/middleware"
	"github.com/go-chi/chi/v5"
)

func New(r chi.Router, ctx context.Context) {
	r.With(
		middleware.WithChat(ctx),
	).Get(
		"/chats/{chat_id}/messages",
		GetByChat(ctx),
	)

	r.With(
		middleware.WithAgent(ctx),
		middleware.WithBody[CreateBody](ctx, "/messages/create"),
	).Post(
		"/agents/{agent_id}/messages",
		Create(ctx),
	)

	r.With(
		middleware.WithChat(ctx),
		middleware.WithBody[CreateBody](ctx, "/messages/create"),
	).Post(
		"/chats/{chat_id}/messages",
		Create(ctx),
	)
}
