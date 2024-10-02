format:
	go fmt ./...

lint:
	golangci-lint run ./...

build:
	go build -o bin/ginserver ./cmd/gin_server/main.go
	go build -o bin/httpserver ./cmd/http_server/main.go

docker:
	docker compose up

all: format lint build docker