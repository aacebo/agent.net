module github.com/aacebo/agent.net/worker

go 1.23.1

require (
	github.com/aacebo/agent.net/amqp v0.0.0
	github.com/aacebo/agent.net/core v0.0.0
	github.com/aacebo/agent.net/postgres v0.0.0
	github.com/rabbitmq/amqp091-go v1.10.0
)

require (
	github.com/golang-migrate/migrate/v4 v4.17.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)

replace github.com/aacebo/agent.net/core => ../core

replace github.com/aacebo/agent.net/amqp => ../amqp

replace github.com/aacebo/agent.net/postgres => ../postgres
