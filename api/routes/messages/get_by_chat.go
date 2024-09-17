package messages

import (
	"context"
	"net/http"

	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/go-chi/render"
)

func GetByChat(ctx context.Context) http.HandlerFunc {
	messages := ctx.Value("repos.messages").(repos.IMessagesRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		chat := r.Context().Value("chat").(models.Chat)
		messages := messages.GetByChatID(chat.ID)

		render.JSON(w, r, messages)
	}
}
