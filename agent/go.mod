module github.com/aacebo/agent.net/agent

go 1.23.1

require (
	github.com/aacebo/agent.net/core v0.0.0
	github.com/aacebo/agent.net/ws v0.0.0
	github.com/go-chi/chi/v5 v5.1.0
	github.com/go-chi/cors v1.2.1
	github.com/go-chi/httprate v0.14.1
	github.com/go-chi/render v1.0.3
	github.com/gorilla/websocket v1.5.3
)

require (
	github.com/ajg/form v1.5.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
)

replace github.com/aacebo/agent.net/core => ../core

replace github.com/aacebo/agent.net/ws => ../ws
