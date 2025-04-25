package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/rezaabaskhanian/toDoList_khodam/internal/domain"
	Storage "github.com/rezaabaskhanian/toDoList_khodam/internal/infrastructure"
	Cli "github.com/rezaabaskhanian/toDoList_khodam/internal/presentation"
	Task "github.com/rezaabaskhanian/toDoList_khodam/internal/usecase/task"
)

func main() {
	// ساختن لایه‌های برنامه
	repo := Storage.NewFileTaskRepository("tasks.json") // ذخیره‌سازی در فایل
	service := Task.NewTaskService(repo)                // ساختن سرویس با ریپو
	cliApp := Cli.NewTaskCli(*service)                  // ساختن CLI با سرویس

	// ایجاد scanner برای خواندن ورودی از کاربر
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// نمایش منو
		fmt.Println("\nلطفاً یکی از گزینه‌ها را انتخاب کنید:")
		fmt.Println("1. ایجاد تسک جدید")
		fmt.Println("2. نمایش لیست تسک‌ها")
		fmt.Println("3. حذف تسک")
		fmt.Println("4. علامت زدن تسک به عنوان انجام‌شده")
		fmt.Println("5. خروج")

		// گرفتن ورودی از کاربر
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			// ایجاد تسک جدید
			fmt.Print("عنوان تسک را وارد کنید: ")
			scanner.Scan()
			title := scanner.Text()

			fmt.Print("توضیحات تسک را وارد کنید: ")
			scanner.Scan()
			desc := scanner.Text()

			// ساختن تسک جدید
			task := domain.Task{
				Title:       title,
				Description: desc,
			}

			// ایجاد تسک از طریق CLI
			cliApp.CreateTask(task)

		case "2":
			// نمایش لیست تسک‌ها
			cliApp.ListTasks()

		case "3":
			// حذف تسک
			fmt.Print("ID تسک را وارد کنید: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("خطا در وارد کردن ID")
				continue
			}

			// حذف تسک
			cliApp.DeleteTask(id)

		case "4":
			// علامت زدن تسک به عنوان انجام‌شده
			fmt.Print("ID تسک را وارد کنید: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("خطا در وارد کردن ID")
				continue
			}

			// علامت زدن تسک به عنوان انجام‌شده
			cliApp.MarkAsDone(id)

		case "5":
			// خروج از برنامه
			fmt.Println("خروج از برنامه. خداحافظ!")
			return

		default:
			// ورودی نامعتبر
			fmt.Println("گزینه نامعتبر است. دوباره تلاش کنید.")
		}
	}
}
