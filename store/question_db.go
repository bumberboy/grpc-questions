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

func (db *QuestionDB) Create(q model.Question) (*model.Question, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	newID := int64(len(db.store))
	q.ID = newID
	q.CreatedAt = time.Now()
	q.UpdatedAt = time.Now()
	db.store[newID] = q
	return &q, nil
}

func (db *QuestionDB) Update(q model.Question) (*model.Question, error) {
	old, err := db.Read(q.ID)
	if err != nil {
		return nil, err
	}
	db.mu.Lock()
	defer db.mu.Unlock()
	q.CreatedAt = old.CreatedAt
	q.UpdatedAt = time.Now()
	db.store[q.ID] = q
	return &q, nil
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

func (db *QuestionDB) Read(id int64) (*model.Question, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	q, ok := db.store[id]
	if !ok {
		return nil, fmt.Errorf("question %d not found", id)
	}
	return &q, nil
}
