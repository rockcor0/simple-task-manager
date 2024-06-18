package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name        string
	description string
	done        bool
}

type TaskList struct {
	tasks []Task
}

// Add a task to the task list
func (taskList *TaskList) addTask(task Task) {
	taskList.tasks = append(taskList.tasks, task)
}

// Mark a task done
func (taskList *TaskList) setDone(index int) {
	taskList.tasks[index].done = true
}

// Edit a task
func (taskList *TaskList) editTask(index int, task Task) {
	taskList.tasks[index] = task
}

// Remove a task
func (taskList *TaskList) removeTask(index int) {
	taskList.tasks = append(taskList.tasks[:index], taskList.tasks[index+1:]...)
}

func main() {
	taskList := TaskList{}
	printMenu(taskList)
	printTasks(taskList)

}

func printMenu(taskList TaskList) {
	read := bufio.NewReader(os.Stdin)
	for {
		var option int
		fmt.Println("Select an option:\n",
			"1. Add a new task\n",
			"2. Set task like done\n",
			"3. Edit a task\n",
			"4. Delete a task\n",
			"5. Exit")
		fmt.Print("Select an option: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			var task Task
			fmt.Print("Enter the name of the task: ")
			task.name, _ = read.ReadString('\n')
			fmt.Print("Enter the description of the task: ")
			task.description, _ = read.ReadString('\n')
			taskList.addTask(task)
			fmt.Println("Task added succesfully")
		case 2:
			var index int
			fmt.Print("Enter the number of the task to done: ")
			fmt.Scan(&index)
			taskList.setDone(index)
			fmt.Println("Task set as done successfully")
		case 3:
			var index int
			fmt.Print("Enter the number of the task to edit: ")
			fmt.Scan(&index)
			var task Task
			fmt.Print("Enter the name of the task: ")
			task.name, _ = read.ReadString('\n')
			fmt.Print("Enter the description of the task: ")
			task.description, _ = read.ReadString('\n')
			taskList.editTask(index, task)
		case 4:
			var index int
			fmt.Print("Enter the number of the task to delete: ")
			fmt.Scan(&index)
			taskList.removeTask(index)
			fmt.Println("Task removed successfully.")
		case 5:
			fmt.Println("Exit...")
			return
		default:
			fmt.Println("Ingrese una opción válida.")
		}
		printTasks(taskList)
	}
}

func printTasks(taskList TaskList) {
	fmt.Println("Task list")
	fmt.Println("==========================================================")
	for i, tasks := range taskList.tasks {
		fmt.Printf("%d. %s - %s - Done: %t\n", i, tasks.name, tasks.description, tasks.done)
	}
	fmt.Println("==========================================================")
}
