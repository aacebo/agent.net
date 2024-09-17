package middleware

import (
	"context"
	"net/http"

	"github.com/aacebo/agent.net/core/repos"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func WithChat(ctx context.Context) func(http.Handler) http.Handler {
	agents := ctx.Value("repos.agents").(repos.IAgentsRepository)
	chats := ctx.Value("repos.chats").(repos.IChatsRepository)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			id := chi.URLParam(r, "chat_id")

			if _, err := uuid.Parse(id); err != nil {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, "invalid `chat_id`")
				return
			}

			chat, exists := chats.GetByID(id)

			if !exists {
				render.Status(r, http.StatusNotFound)
				render.JSON(w, r, "chat not found")
				return
			}

			agent, exists := agents.GetByID(chat.AgentID)

			if !exists {
				render.Status(r, http.StatusNotFound)
				render.JSON(w, r, "agent not found")
				return
			}

			ctx = context.WithValue(ctx, "agent", agent)
			ctx = context.WithValue(ctx, "chat", chat)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
