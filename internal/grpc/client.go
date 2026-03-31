package grpcclient

import (
    "context"
    "time"

    pb "github.com/DennisMRitchie/go-rag-llm-orchestrator/proto/llm"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

type LLMClient struct {
    conn   *grpc.ClientConn
    client pb.LLMServiceClient
}

func NewLLMClient(address string) (*LLMClient, error) {
    conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        return nil, err
    }
    return &LLMClient{
        conn:   conn,
        client: pb.NewLLMServiceClient(conn),
    }, nil
}

func (c *LLMClient) Generate(ctx context.Context, prompt string, contextChunks []string, temp float32) (string, float32, error) {
    ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
    defer cancel()

    req := &pb.GenerateRequest{
        Prompt:        prompt,
        ContextChunks: contextChunks,
        Temperature:   temp,
    }

    resp, err := c.client.Generate(ctx, req)
    if err != nil {
        return "", 0, err
    }
    return resp.Answer, resp.Confidence, nil
}

func (c *LLMClient) Close() {
    c.conn.Close()
}
