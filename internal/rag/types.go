package rag

type Document struct {
    ID       string            `json:"id"`
    Content  string            `json:"content"`
    Metadata map[string]string `json:"metadata"`
}

type Chunk struct {
    ID        string    `json:"id"`
    DocID     string    `json:"doc_id"`
    Content   string    `json:"content"`
    Embedding []float32 `json:"embedding"`
}

type QueryRequest struct {
    Query       string  `json:"query"`
    TopK        int     `json:"top_k"`
    Temperature float32 `json:"temperature"`
}

type QueryResponse struct {
    Answer  string   `json:"answer"`
    Sources []string `json:"sources"`
    Latency float64  `json:"latency_ms"`
}
