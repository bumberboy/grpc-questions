package question

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	questionpb "grpc-questions/gen/go/proto"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	RegisterService(s, "", nil)
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func createTestClient() (questionpb.QuestionServiceClient, *grpc.ClientConn) {
	ctx := context.Background()
	opt := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), opt)
	if err != nil {
		panic(err)
	}
	client := questionpb.NewQuestionServiceClient(conn)
	return client, conn
}

func TestCreateQuestion(t *testing.T) {
	client, conn := createTestClient()
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &questionpb.CreateQuestionRequest{
		Question: &questionpb.Question{
			Title:       "What is Golang?",
			Body:        "What is Golang and its features?",
			Answer:      "Golang is a programming language...",
			Explanation: "Golang, also known as Go, is a statically typed, compiled language...",
			Params:      map[string]string{"difficulty": "easy"},
		},
	}

	resp, err := client.CreateQuestion(ctx, request)
	if err != nil {
		t.Fatalf("Failed to create question: %v", err)
	}
	if resp.Question.Title != request.Question.Title {
		t.Errorf("Expected title %s, got %s", request.Question.Title, resp.Question.Title)
	}
}

func TestReadQuestion(t *testing.T) {
	client, conn := createTestClient()
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a question to read later
	createRequest := &questionpb.CreateQuestionRequest{
		Question: &questionpb.Question{
			Title:       "What is Golang?",
			Body:        "What is Golang and its features?",
			Answer:      "Golang is a programming language...",
			Explanation: "Golang, also known as Go, is a statically typed, compiled language...",
			Params:      map[string]string{"difficulty": "easy"},
		},
	}
	createResp, err := client.CreateQuestion(ctx, createRequest)
	if err != nil {
		t.Fatalf("Failed to create question: %v", err)
	}

	// Read the created question
	readRequest := &questionpb.ReadQuestionRequest{
		Id: createResp.Question.Id,
	}
	readResp, err := client.ReadQuestion(ctx, readRequest)
	if err != nil {
		t.Fatalf("Failed to read question: %v", err)
	}
	if readResp.Question.Title != createRequest.Question.Title {
		t.Errorf("Expected title %s, got %s", createRequest.Question.Title, readResp.Question.Title)
	}
}

func TestUpdateQuestion(t *testing.T) {
	client, conn := createTestClient()
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a question to update later
	createRequest := &questionpb.CreateQuestionRequest{
		Question: &questionpb.Question{
			Title:       "What is Golang?",
			Body:        "What is Golang and its features?",
			Answer:      "Golang is a programming language...",
			Explanation: "Golang, also known as Go, is a statically typed, compiled language...",
			Params:      map[string]string{"difficulty": "easy"},
		},
	}
	createResp, err := client.CreateQuestion(ctx, createRequest)
	if err != nil {
		t.Fatalf("Failed to create question: %v", err)
	}

	// Update the created question
	updateRequest := &questionpb.UpdateQuestionRequest{
		Question: &questionpb.Question{
			Id:          createResp.Question.Id,
			Title:       "What is Go programming language?",
			Body:        "What is Go programming language and its features?",
			Answer:      "Go is a programming language...",
			Explanation: "Go, also known as Golang, is a statically typed, compiled language...",
			Params:      map[string]string{"difficulty": "medium"},
		},
	}
	updateResp, err := client.UpdateQuestion(ctx, updateRequest)
	if err != nil {
		t.Fatalf("Failed to update question: %v", err)
	}
	if updateResp.Question.Title != updateRequest.Question.Title {
		t.Errorf("Expected title %s, got %s", updateRequest.Question.Title, updateResp.Question.Title)
	}
}

func TestListQuestions(t *testing.T) {
	client, conn := createTestClient()
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a few questions to list later
	questions := []*questionpb.Question{
		{
			Title:       "What is Golang?",
			Body:        "What is Golang and its features?",
			Answer:      "Golang is a programming language...",
			Explanation: "Golang, also known as Go, is a statically typed, compiled language...",
			Params:      map[string]string{"difficulty": "easy"},
		},
		{
			Title:       "What is Python?",
			Body:        "What is Python and its features?",
			Answer:      "Python is a programming language...",
			Explanation: "Python is a high-level, interpreted, and dynamically typed programming language...",
			Params:      map[string]string{"difficulty": "easy"},
		},
	}

	for _, question := range questions {
		request := &questionpb.CreateQuestionRequest{
			Question: question,
		}
		_, err := client.CreateQuestion(ctx, request)
		if err != nil {
			t.Fatalf("Failed to create question: %v", err)
		}
	}

	// List questions
	listRequest := &questionpb.ListQuestionsRequest{}
	listResp, err := client.ListQuestions(ctx, listRequest)
	if err != nil {
		t.Fatalf("Failed to list questions: %v", err)
	}

	// Verify if the number of questions returned is at least as many as the created ones
	if len(listResp.Questions) < len(questions) {
		t.Errorf("Expected at least %d questions, got %d", len(questions), len(listResp.Questions))
	}
}
