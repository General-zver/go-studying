package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	id          int
	title       string
	description string
	fileName    string = "improvised_base.txt"
	reader      bufio.Reader
	scanner     bufio.Scanner
)

func main() {
	reader = *bufio.NewReader(os.Stdin)
	scanner = *bufio.NewScanner(os.Stdin)
	fmt.Println("Доступные команды\n", "new\n", "update\n", "delete\n", "search\n", "exit\n")

	for {
		fmt.Print(">> ")

		if !scanner.Scan() {
			break
		}

		command := strings.TrimSpace(scanner.Text())
		switch command {
		case "new":
			for {
				scanDataFromConsole("Введите название задачи: ", &title)
				if len(title) == 0 {
					fmt.Println("Название не может быть пустым")
					scanDataFromConsole("Введите название задачи: ", &title)
				} else {
					break
				}
			}
			scanDataFromConsole("Введите описание задачи: ", &description)

			//fmt.Println(task.Title)

			opened_file := openFileOrCreate(fileName)
			existing_data := readDataFromFile(opened_file)

			tasks := decodeFromJson(existing_data)
			id = getLastIdFromTasks(tasks) + 1
			task := createTask(id, title, description)

			tasks = append(tasks, task)

			prepared_task := prepareTaskForWrite(tasks)
			// fmt.Println("tasks prepared")
			// fmt.Println(string(prepared_task))

			if err := file.Truncate(0); err != nil {
				panic(err)
			}
			if _, err := file.Seek(0, io.SeekStart); err != nil {
				panic(err)
			}

			opened_file.Write(prepared_task)

			fmt.Println(string("Задача Зафиксирована"))

			opened_file.Close()
		case "update":
			var update_id string
			opened_file := openFileOrCreate(fileName)
			existing_data := readDataFromFile(opened_file)

			tasks := decodeFromJson(existing_data)
			for _, v := range tasks {
				fmt.Printf("%d. %s", v.Id, v.Title)
				fmt.Printf("->%s\n", v.Description)
			}
			scanDataFromConsole("Введите номер задачи: ", &update_id)

			scanDataFromConsole("Введите название задачи: ", &title)
			scanDataFromConsole("Введите описание задачи: ", &description)

			ch_t := changeTask(getTaskById(update_id, tasks), title, description)
			fmt.Println("Задача обновлена:")
			fmt.Printf("%d. %s", ch_t.Id, ch_t.Title)
			fmt.Printf("->%s\n", ch_t.Description)

		case "delete":
			fmt.Println("Данный функционал еще не реализован.")
		case "search":

			fmt.Println("Данный функционал еще не реализован.")
		case "show all tasks":
			opened_file := openFileOrCreate(fileName)
			existing_data := readDataFromFile(opened_file)

			tasks := decodeFromJson(existing_data)
			for _, v := range tasks {
				fmt.Printf("%d. %s", v.Id, v.Title)
				fmt.Printf("->%s\n", v.Description)
			}
			// fmt.Println(string(existing_data))

			opened_file.Close()
		case "exit":
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

func scanDataFromConsole(message string, variable *string) {
	var err error
	fmt.Print(message)
	*variable, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Ошибка чтения строки:", err)
	} else {
		fmt.Println("Вы ввели:", *variable)
	}
}
