package domain

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatAt     time.Time `json:"creatAt"`
	Done        bool      `json:"done"`
}
