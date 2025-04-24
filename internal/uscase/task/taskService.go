package Task

import (
	"fmt"
	"time"

	"github.com/rezaabaskhanian/toDoList_khodam/internal/domain"
)

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (u *TaskService) CreateTask(title, description string) (domain.Task, error) {

	tasks, err := u.repo.Load()

	if err != nil {
		return domain.Task{}, err
	}

	newID := 1

	for _, t := range tasks {
		if t.ID >= newID {
			newID = t.ID + 1
		}
	}

	task := domain.Task{
		ID:          newID,
		Title:       title,
		Description: description,
		CreatAt:     time.Now(),
		Done:        false,
	}
	tasks = append(tasks, task)

	if err := u.repo.Save(tasks); err != nil {
		return domain.Task{}, err
	}

	return task, nil

}

// list of tasks
func (u *TaskService) ListTasks() ([]domain.Task, error) {
	tasks, err := u.repo.Load()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// user can change tasks to done true

func (u *TaskService) MarkAsDone(taskId int) (domain.Task, error) {
	tasks, err := u.repo.Load()
	if err != nil {
		return domain.Task{}, err
	}

	found := false
	var newTask domain.Task

	for i, t := range tasks {

		if taskId == t.ID {
			tasks[i].Done = true
			newTask = tasks[i]
			found = true
			break
		}

	}
	if !found {
		fmt.Println("همچین تسکی با این آیدی وجود ندارد")
	}

	err = u.repo.Save(tasks)
	if err != nil {
		return domain.Task{}, err
	}
	fmt.Println("تسک اوکی شد ")
	return newTask, nil

}

// user can delete user by id
func (u *TaskService) DeleteById(taskId int) error {
	tasks, err := u.repo.Load()

	if err != nil {
		return err
	}

	found := false
	var result []domain.Task
	for _, t := range tasks {
		if t.ID == taskId {
			found = true

		} else {
			result = append(result, t)
		}

	}
	if !found {
		fmt.Println("همچین تسکی با این آیدی وجود ندارد")
	}
	return u.repo.Save(result)

}
