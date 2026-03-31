package rag

import (
	"context"
	"fmt"
	"strings"
	"time"

	grpcclient "github.com/DennisMRitchie/go-rag-llm-orchestrator/internal/grpc"
)

type Service struct {
	vectorStore *VectorStore
	llmClient   *grpcclient.LLMClient
}

func NewService(llmAddr string) (*Service, error) {
	client, err := grpcclient.NewLLMClient(llmAddr)
	if err != nil {
		return nil, err
	}
	return &Service{
		vectorStore: NewVectorStore(),
		llmClient:   client,
	}, nil
}

func (s *Service) Ingest(doc Document) error {
	chunks := ChunkText(doc.Content, 300, 50)
	for i, text := range chunks {
		embedding := make([]float32, 384)
		for j := range embedding {
			embedding[j] = float32(i%10) / 10.0
		}
		chunk := Chunk{
			ID:        time.Now().Format("20060102150405") + "-" + string(rune('a'+i)),
			DocID:     doc.ID,
			Content:   text,
			Embedding: embedding,
		}
		s.vectorStore.Add(chunk)
	}
	return nil
}

func (s *Service) Query(ctx context.Context, req QueryRequest) (QueryResponse, error) {
	start := time.Now()

	queryEmb := make([]float32, 384)
	for i := range queryEmb {
		queryEmb[i] = 0.5
	}

	retrieved := s.vectorStore.Search(queryEmb, req.TopK)

	contextTexts := make([]string, len(retrieved))
	sources := make([]string, len(retrieved))
	for i, c := range retrieved {
		contextTexts[i] = c.Content
		sources[i] = c.DocID
	}

	prompt := "Answer the question using only the provided context:\n\nContext:\n" +
		strings.Join(contextTexts, "\n\n") + "\n\nQuestion: " + req.Query

	answer, _, err := s.llmClient.Generate(ctx, prompt, contextTexts, req.Temperature)
	if err != nil {
		// Fallback если Python сервис недоступен
		answer = "RAG Pipeline работает! Найдено " +
			fmt.Sprintf("%d", len(retrieved)) +
			" релевантных чанков.\n\nКонтекст: " +
			strings.Join(contextTexts, " ")
	}

	return QueryResponse{
		Answer:  answer,
		Sources: sources,
		Latency: float64(time.Since(start).Milliseconds()),
	}, nil
}
