package Cli

import (
	"fmt"

	"github.com/rezaabaskhanian/toDoList_khodam/internal/domain"

	Task "github.com/rezaabaskhanian/toDoList_khodam/internal/usecase/task"
)

type CliService struct {
	taskCliService Task.TaskService
}

func NewTaskCli(taskCliService Task.TaskService) TaskServiceCli {
	return &CliService{taskCliService: taskCliService}
}

// CreateTask implements TaskServiceCli.
func (c *CliService) CreateTask(task domain.Task) error {
	// فراخوانی سرویس برای ایجاد تسک جدید
	task, err := c.taskCliService.CreateTask(task.Title, task.Description)

	if err != nil {
		// اگر خطایی پیش بیاید، پیام خطا چاپ می‌شود
		fmt.Println("خطا در ایجاد تسک:", err)
		return err
	}
	// اگر تسک با موفقیت ایجاد شود، پیامی به کاربر نمایش داده می‌شود
	fmt.Println("تسک با موفقیت ایجاد شد.", task)
	return nil

}

// DeleteTask implements TaskServiceCli.
func (c *CliService) DeleteTask(id int) error {
	err := c.taskCliService.DeleteById(id)
	if err != nil {
		// اگر خطا پیش آید، پیام خطا چاپ می‌شود
		fmt.Println("خطا در حذف تسک:", err)
		return err
	}
	// اگر تسک با موفقیت حذف شود، پیامی به کاربر نمایش داده می‌شود
	fmt.Println("تسک با موفقیت حذف شد.")
	return nil
}

// ListTasks implements TaskServiceCli.
func (c *CliService) ListTasks() ([]domain.Task, error) {
	// فراخوانی سرویس برای دریافت لیست تسک‌ها
	tasks, err := c.taskCliService.ListTasks()
	if err != nil {
		// اگر خطا پیش آید، پیام خطا چاپ می‌شود
		fmt.Println("خطا در دریافت لیست تسک‌ها:", err)
		return nil, err
	}
	// نمایش تسک‌ها به صورت جدول
	for _, t := range tasks {
		// هر تسک شامل ID، عنوان و وضعیت (انجام‌شده یا خیر) نمایش داده می‌شود
		fmt.Printf("ID: %d | Title: %s | Done: %v\n", t.ID, t.Title, t.Done)
	}
	return tasks, nil
}

// MarkAsDone implements TaskServiceCli.
func (c *CliService) MarkAsDone(id int) (domain.Task, error) {
	task, err := c.taskCliService.MarkAsDone(id)
	if err != nil {
		// اگر خطا پیش آید، پیام خطا چاپ می‌شود
		fmt.Println("خطا در انجام تسک:", err)
		return domain.Task{}, err
	}
	// نمایش پیامی به کاربر که تسک با موفقیت انجام شده است
	fmt.Println("تسک انجام شد:", task.Title)
	return task, nil
}
