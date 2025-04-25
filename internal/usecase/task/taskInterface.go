package Task

import "github.com/rezaabaskhanian/toDoList_khodam/internal/domain"

type TaskRepository interface {
	// برای این یک لیست از آرایه میگیره چون وقتی میخاهیم ذخیره کنیم در دیتابیس کل دیتابیس رو میگیریم
	// که آخرین شماره دیتابیس را بداینم بریا اینکه تسک بعدی را شماره بدیم
	Save(tasks []domain.Task) error

	Load() ([]domain.Task, error)
}
