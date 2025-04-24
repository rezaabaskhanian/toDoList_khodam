package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rezaabaskhanian/toDoList_khodam/internal/domain"
	Storage "github.com/rezaabaskhanian/toDoList_khodam/internal/infrastructure"
	cli "github.com/rezaabaskhanian/toDoList_khodam/internal/presentation"
	"github.com/rezaabaskhanian/toDoList_khodam/internal/usecase/task"
)

func main() {
	// ساختن لایه‌های برنامه
	repo := Storage.NewFileTaskRepository("tasks.json") // ذخیره‌سازی در فایل
	service := task.NewTaskService(repo)                // ساختن سرویس با ریپو
	cliApp := cli.NewTaskCli(service)                   // ساختن CLI با سرویس

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nلطفاً یکی از گزینه‌ها را انتخاب کنید:")
		fmt.Println("1. ایجاد تسک جدید")
		fmt.Println("2. نمایش لیست تسک‌ها")
		fmt.Println("3. حذف تسک")
		fmt.Println("4. علامت زدن تسک به عنوان انجام‌شده")
		fmt.Println("5. خروج")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("عنوان تسک را وارد کنید: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("توضیحات تسک را وارد کنید: ")
			desc, _ := reader.ReadString('\n')
			desc = strings.TrimSpace(desc)

			task := domain.Task{
				Title:       title,
				Description: desc,
			}

			cliApp.CreateTask(task)

		case "2":
			cliApp.ListTasks()

		case "3":
			fmt.Print("ID تسک را وارد کنید: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)

			cliApp.DeleteTask(id)

		case "4":
			fmt.Print("ID تسک را وارد کنید: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)

			cliApp.MarkAsDone(id)

		case "5":
			fmt.Println("خروج از برنامه. خداحافظ!")
			return

		default:
			fmt.Println("گزینه نامعتبر است. دوباره تلاش کنید.")
		}
	}
}
