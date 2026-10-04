package main

import (
	"fmt"
	"taskmanager/task"
)

func main() {
	newTask1 := task.NewTask("Сделать уроки")
	newTask2 := task.NewTask("Пойти в школу")
	newTask3 := task.NewTask("Пойти за хлебом")

	newTask2.Done = true

	tasks := []task.Task{newTask1, newTask2, newTask3}

	for _, v := range tasks {
		status := "[ ]"

		if v.Done == true {
			status = "[X]"
		}

		fmt.Printf("%s %s\n", status, v.Title)

	}
}
