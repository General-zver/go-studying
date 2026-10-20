package main

func getLastIdFromTasks(tasks []task) int {
	var max int = 0
	for _, value := range tasks {
		if value.Id > max {
			max = value.Id
		}
	}
	return max
}
