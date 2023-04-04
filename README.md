## GRPC-Questions

This is a simple question service for a question bank and I'm using this exercise to experiment with creating a HTTP web server using gRPC.

### Tool(s)

- The gRPC-Gateway creates a HTTP server that will proxy requests to the gRPC server.
- `buf` handles the generation of gRPC code, gateway code and swagger spec from the proto file.
- `buf`: https://github.com/bufbuild/buf

### Quick Demo
1. Start server with `go run ./cmd/server/main.go`. This will start a gRPC server on port 8090 and a HTTP server (gRPC gateway) on port 8091.
2. Run curl command to create a question:
```
curl -X POST "http://localhost:8091/v1/questions" -H "Content-Type: application/json" -d '{
  "question": {
    "title": "Tommy has 5 apples",
    "body": "Tommy has 5 apples. He gives 3 to his sister. How many apples does Tommy have left?",
    "answer": "2",
    "explanation": "5 minus 3 is 2",
    "params": {
      "difficulty": "easy"
    }
  }
}'
```
3. Run curl command to list questions:
```
curl -X GET "http://localhost:8091/v1/questions" -H "Content-Type: application/json"
```