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

func changeTask(t_id string, all_tasks []task, title string, description string) task {
	// t := getTaskById(t_id, all_tasks)
	var t task
	t_iid, err := strconv.Atoi(strings.TrimSpace(t_id))
	if err != nil {
		fmt.Println("ID должен быть числом")
	}
	for index, v := range all_tasks {
		if v.Id == t_iid {
			v.Title = strings.TrimSpace(title)
			v.Description = strings.TrimSpace(description)
			all_tasks[index] = v
			t = all_tasks[index]
		}
	}
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

func getTaskIndexById(id string, tasks []task) int {
	var index int
	iid, err := strconv.Atoi(strings.TrimSpace(id))
	if err != nil {
		panic(err)
	}
	for i, v := range tasks {
		if v.Id == iid {
			index = i
		}
	}
	return index

}

func searchTaskByWord(word string, tasks []task) []task {
	var resul_tasks_arr []task
	for _, v := range tasks {
		if strings.Contains(strings.TrimSpace(v.Title), strings.TrimSpace(word)) {
			resul_tasks_arr = append(resul_tasks_arr, v)
		} else if strings.Contains(strings.TrimSpace(v.Description), strings.TrimSpace(word)) {
			resul_tasks_arr = append(resul_tasks_arr, v)
		}
	}
	return resul_tasks_arr
}
