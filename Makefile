format:
	go fmt ./...

lint:
	golangci-lint run ./...

build:
	go build -o bin/ginserver ./cmd/ginmain/main.go
	go build -o bin/httpserver ./cmd/httpmain/main.go

docker:
	docker compose up

all: format lint build docker