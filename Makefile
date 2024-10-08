fmt:
	gofmt -w ./

test:
	go clean -testcache
	go test ./... -cover

run: api.run worker.run

api.run:
	cd api ; LOG_PREFIX=agent.net/api make run

worker.run:
	cd worker ; LOG_PREFIX=agent.net/worker make run

migrate.up:
	migrate -source file://postgres/migrations -database "$(POSTGRES_CONNECTION_STRING)" up

migrate.down:
	migrate -source file://postgres/migrations -database "$(POSTGRES_CONNECTION_STRING)" down -all

migrate.new:
	migrate create -ext sql -dir postgres/migrations $(name)
