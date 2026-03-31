package rag

import (
    "math"
    "sync"
)

type VectorStore struct {
    chunks map[string]Chunk
    mu     sync.RWMutex
}

func NewVectorStore() *VectorStore {
    return &VectorStore{chunks: make(map[string]Chunk)}
}

func cosineSimilarity(a, b []float32) float32 {
    if len(a) != len(b) {
        return 0
    }
    var dot, normA, normB float32
    for i := range a {
        dot += a[i] * b[i]
        normA += a[i] * a[i]
        normB += b[i] * b[i]
    }
    if normA == 0 || normB == 0 {
        return 0
    }
    return dot / (float32(math.Sqrt(float64(normA))) * float32(math.Sqrt(float64(normB))))
}

func (vs *VectorStore) Add(chunk Chunk) {
    vs.mu.Lock()
    defer vs.mu.Unlock()
    vs.chunks[chunk.ID] = chunk
}

func (vs *VectorStore) Search(queryEmbedding []float32, topK int) []Chunk {
    vs.mu.RLock()
    defer vs.mu.RUnlock()

    type scored struct {
        chunk Chunk
        score float32
    }

    scores := make([]scored, 0, len(vs.chunks))
    for _, c := range vs.chunks {
        score := cosineSimilarity(queryEmbedding, c.Embedding)
        scores = append(scores, scored{chunk: c, score: score})
    }

    for i := 0; i < len(scores)-1; i++ {
        for j := i + 1; j < len(scores); j++ {
            if scores[i].score < scores[j].score {
                scores[i], scores[j] = scores[j], scores[i]
            }
        }
    }

    result := make([]Chunk, 0, topK)
    for i := 0; i < topK && i < len(scores); i++ {
        result = append(result, scores[i].chunk)
    }
    return result
}
