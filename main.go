package main

import (
    "context"
    "log"
    "net/http"
    "os"

    "github.com/DennisMRitchie/go-rag-llm-orchestrator/internal/rag"
    "github.com/gin-gonic/gin"
)

func main() {
    llmAddr := os.Getenv("LLM_SERVICE_ADDR")
    if llmAddr == "" {
        llmAddr = "localhost:50051"
    }

    service, err := rag.NewService(llmAddr)
    if err != nil {
        log.Fatalf("Failed to create RAG service: %v", err)
    }

    _ = service.Ingest(rag.Document{
        ID:      "demo-doc-1",
        Content: "Go is a statically typed compiled programming language designed at Google. It is known for its excellent concurrency support with goroutines and channels. Go is widely used for building microservices, cloud infrastructure, and high-performance APIs.",
        Metadata: map[string]string{"source": "wikipedia"},
    })

    _ = service.Ingest(rag.Document{
        ID:      "demo-doc-2",
        Content: "RAG (Retrieval-Augmented Generation) is a technique that combines information retrieval with language generation. It retrieves relevant documents from a knowledge base and uses them as context for generating accurate answers.",
        Metadata: map[string]string{"source": "research-paper"},
    })

    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    r.POST("/query", func(c *gin.Context) {
        var req rag.QueryRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if req.TopK == 0 {
            req.TopK = 3
        }
        if req.Temperature == 0 {
            req.Temperature = 0.7
        }

        resp, err := service.Query(context.Background(), req)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, resp)
    })

    log.Println("Go RAG Orchestrator running on :8080")
    log.Fatal(r.Run(":8080"))
}
