module github.com/aacebo/agent.net/postgres

go 1.23.1

require (
	github.com/aacebo/agent.net/core v0.0.0
	github.com/golang-migrate/migrate/v4 v4.17.1
	github.com/lib/pq v1.10.9
)

require (
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)

replace github.com/aacebo/agent.net/core => ../core
