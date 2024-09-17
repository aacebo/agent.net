package messages

import (
	"context"

	"github.com/aacebo/agent.net/amqp"
)

func New(ctx context.Context) {
	amqp := ctx.Value("amqp").(*amqp.Client)
	amqp.Consume("messages", "create", Create(ctx))
}
