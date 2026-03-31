# 🧠 Go RAG LLM Orchestrator

> Production-ready Retrieval-Augmented Generation (RAG) orchestrator built in Go with gRPC LLM integration.

![Go](https://img.shields.io/badge/Go-1.23-00ADD8?style=flat&logo=go)
![Gin](https://img.shields.io/badge/Gin-v1.10-00ACD7?style=flat)
![gRPC](https://img.shields.io/badge/gRPC-ready-4285F4?style=flat&logo=google)
![Docker](https://img.shields.io/badge/Docker-ready-2496ED?style=flat&logo=docker)

## ✨ Features

- ⚡ High-concurrency Go service with Gin framework
- 🔍 In-memory vector store with cosine similarity search
- 📄 Smart text chunking with configurable overlap
- 🤖 gRPC client for Python LLM backend (Ollama/vLLM ready)
- 🔌 Easy to swap vector store with Weaviate or Pinecone
- 🐳 Fully Dockerized with docker-compose
- 🛡️ Graceful error handling with LLM fallback mode

## 🏗️ Architecture
```
┌─────────────────┐     POST /query      ┌──────────────────────┐
│   API Client    │ ──────────────────▶  │   Go RAG Orchestrator │
│   (curl/app)    │ ◀──────────────────  │   (Gin + gRPC)        │
└─────────────────┘                      ├──────────────────────┤
                                         │   Vector Store        │
                                         │   (cosine similarity) │
                                         ├──────────────────────┤
                                         │   Text Chunker        │
                                         ├──────────────────────┤
                                         │   gRPC LLM Client     │
                                         └──────────┬───────────┘
                                                    │
                                         ┌──────────▼───────────┐
                                         │  Python LLM Service   │
                                         │  (Ollama/vLLM/HF)     │
                                         └──────────────────────┘
```

## 🚀 Quick Start

### Local Development
```bash
# Clone the repo
git clone https://github.com/DennisMRitchie/go-rag-llm-orchestrator.git
cd go-rag-llm-orchestrator

# Download dependencies
go mod tidy

# Run the service
go run main.go
```

Server starts on `http://localhost:8080`

### Docker
```bash
make up
```

## 📡 API Endpoints

| Method | Endpoint  | Description        |
|--------|-----------|--------------------|
| GET    | `/health` | Health check       |
| POST   | `/query`  | RAG query endpoint |

### Example Request
```bash
curl -X POST http://localhost:8080/query \
  -H "Content-Type: application/json" \
  -d '{"query": "What is Go programming language?", "top_k": 2}'
```

### Example Response
```json
{
  "answer": "Go is a statically typed compiled programming language designed at Google...",
  "sources": ["demo-doc-1", "demo-doc-2"],
  "latency_ms": 91
}
```

## 📁 Project Structure
```
go-rag-llm-orchestrator/
├── main.go                  # Entry point, REST API handlers
├── internal/
│   ├── rag/
│   │   ├── types.go         # Data models
│   │   ├── vectorstore.go   # In-memory vector store
│   │   ├── chunker.go       # Text chunking logic
│   │   └── service.go       # RAG orchestration logic
│   └── grpc/
│       └── client.go        # gRPC LLM client
├── proto/
│   └── llm.proto            # Protobuf definitions
├── python-llm-service/      # Python LLM companion service
│   ├── app.py
│   ├── requirements.txt
│   └── Dockerfile
├── Dockerfile
├── docker-compose.yml
└── Makefile
```

## 🛠️ Tech Stack

| Layer | Technology |
|-------|-----------|
| Backend | Go 1.23, Gin v1.10 |
| Communication | gRPC, Protocol Buffers |
| Vector Search | In-memory cosine similarity |
| LLM Backend | Python (Ollama/vLLM/HuggingFace) |
| DevOps | Docker, Docker Compose |

## 🔧 Make Commands
```bash
make run     # Run locally
make build   # Build binary
make proto   # Generate proto files
make up      # Start with Docker
make down    # Stop Docker
make test    # Test the API
```

## 🗺️ Roadmap

- [ ] Real embedding model integration (sentence-transformers)
- [ ] Weaviate / Pinecone vector DB support
- [ ] OpenTelemetry tracing
- [ ] Rate limiting middleware
- [ ] Streaming responses

## 📄 License

MIT License

---

Built with ❤️ to demonstrate production-ready Go + LLM/NLP skills for Senior Go Developer roles.