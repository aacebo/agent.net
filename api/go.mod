module github.com/aacebo/agent.net/api

go 1.23.1

require (
	github.com/aacebo/agent.net/amqp v0.0.0
	github.com/aacebo/agent.net/core v0.0.0
	github.com/aacebo/agent.net/postgres v0.0.0
	github.com/aacebo/agent.net/ws v0.0.0
	github.com/go-chi/chi/v5 v5.1.0
	github.com/go-chi/cors v1.2.1
	github.com/go-chi/httprate v0.14.1
	github.com/go-chi/render v1.0.3
	github.com/google/uuid v1.6.0
	github.com/gorilla/websocket v1.5.3
	github.com/santhosh-tekuri/jsonschema/v5 v5.3.1
)

require (
	github.com/ajg/form v1.5.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/golang-migrate/migrate/v4 v4.17.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/rabbitmq/amqp091-go v1.10.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)

replace github.com/aacebo/agent.net/core => ../core

replace github.com/aacebo/agent.net/amqp => ../amqp

replace github.com/aacebo/agent.net/postgres => ../postgres

replace github.com/aacebo/agent.net/ws => ../ws
