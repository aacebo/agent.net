package routes

import (
	"context"

	"github.com/aacebo/agent.net/worker/routes/agents"
	"github.com/aacebo/agent.net/worker/routes/chats"
	"github.com/aacebo/agent.net/worker/routes/messages"
)

func New(ctx context.Context) {
	agents.New(ctx)
	chats.New(ctx)
	messages.New(ctx)
}
