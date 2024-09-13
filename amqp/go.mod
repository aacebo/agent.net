module github.com/aacebo/agent.net/amqp

go 1.23.1

require (
	github.com/aacebo/agent.net/core v0.0.0
	github.com/google/uuid v1.6.0
	github.com/rabbitmq/amqp091-go v1.10.0
)

replace github.com/aacebo/agent.net/core => ../core
