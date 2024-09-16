package agents

import (
	"context"

	"github.com/aacebo/agent.net/amqp"
)

func New(ctx context.Context) {
	amqp := ctx.Value("amqp").(*amqp.Client)
	amqp.Consume("agents", "create", Create(ctx))
	amqp.Consume("agents", "start", Start(ctx))
	amqp.Consume("agents", "stop", Stop(ctx))
}
