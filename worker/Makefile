fmt:
	gofmt -w ./

run:
	go run main.go

build:
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
