
# Variables
APP_NAME := grpc-crud-tutorial


# Commands
rootx:
	go run cmd/rootx/main.go

install:
	go mod tidy

swag:
	swag init -g ./cmd/main.go -o docs

grpc:
	go run cmd/grpc/main.go
	
dev:
	go run cmd/main.go

build:
	go build -o bin/$(APP_NAME) cmd/main.go

run:
	/bin/bash -c "bin/$(APP_NAME)"

deploy:
	docker-compose -f docker-compose.yaml up -d

re-deploy:
	docker-compose -f docker-compose.yaml down
	docker system prune -f
	docker-compose -f docker-compose.yaml up -d --build

down:
	docker-compose -f docker-compose.yaml down
