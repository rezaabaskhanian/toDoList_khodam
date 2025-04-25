package Cli

import "github.com/rezaabaskhanian/toDoList_khodam/internal/domain"

type TaskServiceCli interface {
	CreateTask(task domain.Task) error
	ListTasks() ([]domain.Task, error)
	DeleteTask(id int) error
	MarkAsDone(id int) (domain.Task, error)
}
