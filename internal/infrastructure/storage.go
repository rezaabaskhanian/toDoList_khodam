package Storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rezaabaskhanian/toDoList_khodam/internal/domain"
	Task "github.com/rezaabaskhanian/toDoList_khodam/internal/usecase/task"
)

const taskStorageFile = "task.json"

// type TaskRepository interface {
// 	Save(tasks []domain.Task) error
// 	Load() ([]domain.Task, error)
// }

// FileTaskRepository پیاده‌سازی ذخیره‌سازی تسک‌ها در فایل است.
type FileTaskRepository struct {
	FileName string
	repo     Task.TaskRepository
}

// NewFileTaskRepository سازنده‌ای برای ایجاد ریپازیتوری است.
func NewFileTaskRepository(fileName string) *FileTaskRepository {
	return &FileTaskRepository{FileName: fileName}
}

func (repo *FileTaskRepository) Save(tasks []domain.Task) error {
	file, err := os.OpenFile(repo.FileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("خطا در باز کردن یا ایجاد فایل: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(tasks); err != nil {
		return fmt.Errorf("خطا در کدگذاری JSON: %w", err)
	}
	return nil
}

// Load بارگذاری تسک‌ها از فایل
func (repo *FileTaskRepository) Load() ([]domain.Task, error) {
	file, err := os.Open(repo.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []domain.Task{}, nil
		}
		return nil, fmt.Errorf("خطا در باز کردن فایل: %w", err)
	}
	defer file.Close()

	var tasks []domain.Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, fmt.Errorf("خطا در دیکد کردن JSON: %w", err)
	}
	return tasks, nil
}
