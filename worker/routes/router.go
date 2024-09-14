package routes

import (
	"context"

	"github.com/aacebo/agent.net/worker/routes/agents"
)

func New(ctx context.Context) {
	agents.New(ctx)
}
