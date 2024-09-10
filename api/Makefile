clean:
	rm -rf ./bin

build:
	go build ./...

clean_build: clean build

fmt:
	gofmt -w ./

test:
	go clean -testcache
	go test ./... -cover

run:
	go run ./

migrate.up:
	migrate -source file://postgres/migrations -database "$(POSTGRES_CONNECTION_STRING)" up

migrate.down:
	migrate -source file://postgres/migrations -database "$(POSTGRES_CONNECTION_STRING)" down -all

migrate.new:
	migrate create -ext sql -dir postgres/migrations $(name)
