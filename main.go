package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	id          int
	title       string
	description string
	fileName    string = "improvised_base.txt"
	reader      bufio.Reader
	scanner     bufio.Scanner
	file        *os.File
)

func main() {
	reader = *bufio.NewReader(os.Stdin)
	scanner = *bufio.NewScanner(os.Stdin)
	fmt.Println("Доступные команды\n", "new\n", "update\n", "delete\n", "search\n", "exit\n")

	openFileOrCreate(fileName)

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

			existing_data := readDataFromFile(file)

			tasks := decodeFromJson(existing_data)
			id = getLastIdFromTasks(tasks) + 1
			task := createTask(id, title, description)

			tasks = append(tasks, task)

			writeTasksInFile(tasks)

			fmt.Println(string("Задача Зафиксирована"))

		case "update":
			var update_id string
			existing_data := readDataFromFile(file)

			tasks := decodeFromJson(existing_data)
			for _, v := range tasks {
				fmt.Printf("%d. %s", v.Id, v.Title)
				fmt.Printf("->%s\n", v.Description)
			}
			scanDataFromConsole("Введите номер задачи: ", &update_id)

			scanDataFromConsole("Введите название задачи: ", &title)
			scanDataFromConsole("Введите описание задачи: ", &description)

			ch_t := changeTask(update_id, tasks, title, description)
			fmt.Println("Задача обновлена:")
			fmt.Printf("%d. %s", ch_t.Id, ch_t.Title)
			fmt.Printf("->%s\n", ch_t.Description)

			writeTasksInFile(tasks)

		case "delete":
			var delete_id string
			existing_data := readDataFromFile(file)

			tasks := decodeFromJson(existing_data)
			for _, v := range tasks {
				fmt.Printf("%d. %s", v.Id, v.Title)
				fmt.Printf("->%s\n", v.Description)
			}
			for {
				scanDataFromConsole("Введите номер задачи: ", &delete_id)
				if len(strings.TrimSpace(delete_id)) > 0 {
					break
				} else {
					fmt.Println("Некорректный ввод")
				}
			}
			task_index := getTaskIndexById(delete_id, tasks)
			tasks = append(tasks[:task_index], tasks[task_index+1])

			writeTasksInFile(tasks)

			fmt.Println("Обновленный список. Уадлена задача с id ", delete_id)
			for _, v := range tasks {
				fmt.Printf("%d. %s", v.Id, v.Title)
				fmt.Printf("->%s\n", v.Description)
			}
		case "search":
			var (
				key_word string
				t        task
			)

			existing_data := readDataFromFile(file)

			tasks := decodeFromJson(existing_data)
			for {
				scanDataFromConsole("Введите id или ключевое слово.", &key_word)
				if len(strings.TrimSpace(key_word)) > 0 {
					break
				} else {
					fmt.Println("Некорректный ввод")
				}
			}
			_, err := strconv.Atoi(key_word)
			if err != nil {
				t_arr := searchTaskByWord(key_word, tasks)
				fmt.Println("Найденные совпадения:")
				for _, v := range t_arr {
					fmt.Printf("%d. %s\n", v.Id, v.Title)
					fmt.Printf("->%s\n", v.Description)
				}

			} else {
				t = getTaskById(key_word, tasks)
				fmt.Println("Найденная задача:")
				fmt.Printf("%d. %\n", t.Id, t.Title)
				fmt.Printf("->%s\n", t.Description)
			}

			// fmt.Println("Данный функционал еще не реализован.")
		case "show all tasks":

			existing_data := readDataFromFile(file)

			tasks := decodeFromJson(existing_data)
			for _, v := range tasks {
				fmt.Printf("%d. %s\n", v.Id, v.Title)
				fmt.Printf("->%s\n", v.Description)
			}

		case "exit":
			file.Close()
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

func writeTasksInFile(tasks []task) {
	prepared_task := prepareTaskForWrite(tasks)

	if err := file.Truncate(0); err != nil {
		panic(err)
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		panic(err)
	}

	file.Write(prepared_task)
}
