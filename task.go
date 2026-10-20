package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type task struct {
	Id          int
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
}

func createTask(id int, title string, description string) task {
	return task{
		Id:          id,
		Title:       strings.TrimSpace(title),
		Description: strings.TrimSpace(description),
		Done:        false,
		CreatedAt:   time.Now(),
	}
}

// func deleteTask() {

// }

func changeTask(t task, title string, description string) task {
	t.Title = strings.TrimSpace(title)
	t.Description = strings.TrimSpace(description)
	return t
}

// func filterTask() {

// }

// Преобразование структуры в json (массива)
func prepareTaskForWrite(t []task) []byte {
	json_data, err := json.MarshalIndent(t, "", "	")
	if err != nil {
		fmt.Println("Ошибка преобразования задачи в json: ", err)
		os.Exit(1)
	}
	return json_data
}

// Обратное преобразование из json в объекты
func decodeFromJson(data []byte) []task {
	var tasks []task
	if len(data) > 0 {
		err := json.Unmarshal(data, &tasks)
		if err != nil {
			fmt.Println("Ошибка JSON:", err)
			os.Exit(1)
		}
	}
	return tasks
}

// func searchTask()  {

// }

func getTaskById(id string, tasks []task) task {
	var v_task task
	iid, err := strconv.Atoi(strings.TrimSpace(id))
	if err != nil {
		panic(err)
	}
	for i, v := range tasks {
		if v.Id == iid {
			v_task = tasks[i]
		}
	}
	return v_task
}
