module github.com/aacebo/agent.net/ws

go 1.23.1

require (
	github.com/aacebo/agent.net/core v0.0.0
	github.com/google/uuid v1.6.0
	github.com/gorilla/websocket v1.5.3
)

replace github.com/aacebo/agent.net/core => ../core
