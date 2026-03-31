.PHONY: build run proto up down test

build:
	go build -o bin/rag-orchestrator .

run:
	go run main.go

proto:
	protoc --go_out=. --go-grpc_out=. proto/llm.proto

up:
	docker compose up --build

down:
	docker compose down

test:
	curl -X POST http://localhost:8080/query \
	  -H "Content-Type: application/json" \
	  -d '{"query": "What is Go programming language known for?", "top_k": 2}'
