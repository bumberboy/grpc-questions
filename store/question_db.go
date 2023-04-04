package store

import (
	"fmt"
	"grpc-questions/model"
	"sync"
	"time"
)

type QuestionDB struct {
	mu    sync.RWMutex
	store map[int64]model.Question
}

func NewQuestionDB() *QuestionDB {
	return &QuestionDB{
		store: make(map[int64]model.Question),
	}
}

func (db *QuestionDB) Create(q model.Question) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, exists := db.store[q.ID]; exists {
		return fmt.Errorf("question %d already exists", q.ID)
	}

	q.CreatedAt = time.Now()
	db.store[q.ID] = q
	return nil
}

func (db *QuestionDB) Update(q model.Question) error {
	_, err := db.Read(q.ID)
	if err != nil {
		return err
	}
	db.mu.Lock()
	defer db.mu.Unlock()
	q.UpdatedAt = time.Now()
	db.store[q.ID] = q
	return nil
}

func (db *QuestionDB) List() ([]model.Question, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	var s []model.Question
	for _, v := range db.store {
		s = append(s, v)
	}
	return s, nil
}

func (db *QuestionDB) Read(id int64) (model.Question, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	q, ok := db.store[id]
	if !ok {
		return model.Question{}, fmt.Errorf("question %d not found", id)
	}
	return q, nil
}
