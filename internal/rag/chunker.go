package rag

import "strings"

func ChunkText(text string, chunkSize, overlap int) []string {
    words := strings.Fields(text)
    var chunks []string
    for i := 0; i < len(words); i += chunkSize - overlap {
        end := i + chunkSize
        if end > len(words) {
            end = len(words)
        }
        chunk := strings.Join(words[i:end], " ")
        chunks = append(chunks, chunk)
        if end == len(words) {
            break
        }
    }
    return chunks
}
