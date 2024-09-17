package chats

import (
	"context"

	"github.com/aacebo/agent.net/amqp"
)

func New(ctx context.Context) {
	amqp := ctx.Value("amqp").(*amqp.Client)
	amqp.Consume("chats", "create", Create(ctx))
}
