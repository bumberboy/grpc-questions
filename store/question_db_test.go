package store

import (
	"grpc-questions/model"
	"testing"
)

func TestCreate(t *testing.T) {
	db := NewQuestionDB()
	q := model.Question{
		Title:       "What is Golang?",
		Body:        "What is Golang and its features?",
		Answer:      "Golang is a programming language...",
		Explanation: "Golang, also known as Go, is a statically typed, compiled language...",
		Params:      map[string]string{"difficulty": "easy"},
	}

	created, err := db.Create(q)
	if err != nil {
		t.Fatalf("Create() failed: %v", err)
	}

	if created.Title != q.Title {
		t.Errorf("Expected title %s, got %s", q.Title, created.Title)
	}
}

func TestRead(t *testing.T) {
	db := setupDB()
	read, err := db.Read(0)
	if err != nil {
		t.Fatalf("Read() failed: %v", err)
	}

	if read.Title != "What is Golang?" {
		t.Errorf("Expected title 'What is Golang?', got %s", read.Title)
	}
}

func TestUpdate(t *testing.T) {
	db := setupDB()
	q, _ := db.Read(0)

	updatedQuestion := *q
	updatedQuestion.Title = "What is the Go programming language?"

	updated, err := db.Update(updatedQuestion)
	if err != nil {
		t.Fatalf("Update() failed: %v", err)
	}

	if updated.Title != updatedQuestion.Title {
		t.Errorf("Expected updated title %s, got %s", updatedQuestion.Title, updated.Title)
	}
}

func TestList(t *testing.T) {
	db := setupDB()
	list, err := db.List()
	if err != nil {
		t.Fatalf("List() failed: %v", err)
	}

	if len(list) != 1 {
		t.Errorf("Expected 1 question, got %d", len(list))
	}
}

func TestUpdateNonexistent(t *testing.T) {
	db := setupDB()
	nonexistentQuestion := model.Question{
		ID:          99,
		Title:       "Nonexistent Question",
		Body:        "This question does not exist in the database.",
		Answer:      "N/A",
		Explanation: "N/A",
		Params:      map[string]string{"difficulty": "N/A"},
	}

	_, err := db.Update(nonexistentQuestion)
	if err == nil {
		t.Error("Expected an error for updating a nonexistent question, but got nil")
	}
}

func TestReadNonexistent(t *testing.T) {
	db := setupDB()
	_, err := db.Read(99)
	if err == nil {
		t.Error("Expected an error for a nonexistent question, but got nil")
	}
}

func setupDB() *QuestionDB {
	db := NewQuestionDB()
	q := model.Question{
		Title:       "What is Golang?",
		Body:        "What is Golang and its features?",
		Answer:      "Golang is a programming language...",
		Explanation: "Golang, also known as Go, is a statically typed, compiled language...",
		Params:      map[string]string{"difficulty": "easy"},
	}

	db.Create(q)
	return db
}
