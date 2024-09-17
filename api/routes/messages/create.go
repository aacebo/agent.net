package messages

import (
	"context"
	"net/http"

	"github.com/aacebo/agent.net/amqp"
	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/aacebo/agent.net/ws"
	"github.com/go-chi/render"
)

type CreateBody struct {
	Text string `json:"text"`
}

func Create(ctx context.Context) http.HandlerFunc {
	amqp := ctx.Value("amqp").(*amqp.Client)
	sockets := ctx.Value("sockets").(*ws.Sockets)
	chats := ctx.Value("repos.chats").(repos.IChatsRepository)
	messages := ctx.Value("repos.messages").(repos.IMessagesRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		agent := r.Context().Value("agent").(models.Agent)
		body := r.Context().Value("body").(CreateBody)
		chat, exists := r.Context().Value("chat").(models.Chat)

		if agent.Status != models.AGENT_STATUS_UP {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, "agent is not started")
			return
		}

		if !exists {
			chat := models.NewChat()
			chat.AgentID = agent.ID
			chat = chats.Create(chat)
			amqp.Publish("chats", "create", chat)
		}

		message := models.NewMessage()
		message.ChatID = chat.ID
		message.FromID = agent.ID
		message.Text = body.Text
		message = messages.Create(message)
		amqp.Publish("messages", "create", message)

		sockets.Send(ws.NewTextMessage(message.Text).WithToID(agent.ID))
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, message)
	}
}
