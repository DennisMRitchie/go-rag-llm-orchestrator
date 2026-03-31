import grpc
from concurrent import futures
import sys
import os

sys.path.insert(0, os.path.dirname(__file__))

class LLMServiceServicer:
    def Generate(self, request, context):
        context_str = "\n".join(request.context_chunks)
        answer = (
            f"[LLM Response]\n"
            f"Based on the retrieved context, here is the answer to your question:\n"
            f"'{request.prompt}'\n\n"
            f"Context used:\n{context_str[:300]}\n\n"
            f"This is a simulated response from a Qwen/Llama-style model. "
            f"In production, this would call a real LLM via Ollama or vLLM."
        )
        return {"answer": answer, "confidence": 0.92}

def serve():
    print("Python LLM gRPC service running on :50051")
    import time
    while True:
        time.sleep(1)

if __name__ == '__main__':
    serve()
