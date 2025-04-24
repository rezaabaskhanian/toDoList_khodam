package domain

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatAt     string `json:"creatAt"`
	Done        bool   `json:"done"`
}
