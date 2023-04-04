package model

import "time"

type Question struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Body        string
	Answer      string
	Explanation string
	Params      map[string]string
}
