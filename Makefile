fmt:
	gofmt -w ./

test:
	go clean -testcache
	go test ./... -cover

api.run:
	go run ./api

worker.run:
	go run ./worker

migrate.up:
	migrate -source file://postgres/migrations -database "$(POSTGRES_CONNECTION_STRING)" up

migrate.down:
	migrate -source file://postgres/migrations -database "$(POSTGRES_CONNECTION_STRING)" down -all

migrate.new:
	migrate create -ext sql -dir postgres/migrations $(name)
