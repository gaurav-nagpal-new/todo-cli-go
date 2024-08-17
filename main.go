package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// todo-cli app

	// take user input
	scanner := bufio.NewScanner(os.Stdin)
	tasks := []string{}

	for {
		fmt.Println("-------------------------------------------")
		fmt.Println("Please type either of option, e.g 1,2...")
		fmt.Println("1. Add Task")
		fmt.Println("2. Remove Task")
		fmt.Println("3. Mark Task as Complete")
		fmt.Println("4. List Tasks")
		fmt.Println("5. Exit")

		scanner.Scan()
		if scanner.Text() == "" {
			fmt.Println("You didn't selected anything")
			break
		}

		switch scanner.Text() {
		case "1":
			fmt.Println("Enter the task. If there are more tasks, enter in comma separated string")
			scanner.Scan()
			if scanner.Text() == "" {
				fmt.Println("You have not entered any task. Exiting...")
				continue
			}
			inputString := strings.ReplaceAll(scanner.Text(), " ", "")
			tasks = append(tasks, strings.Split(inputString, ",")...)
			fmt.Println("Task/s Added")
		case "2":
			fmt.Println("Enter the task you want to remove")
			scanner.Scan()
			if scanner.Text() == "" {
				fmt.Println("You have not entered any task. Exiting...")
				continue
			}
			// find the index of the task which is supposed to be removed
			idx := -1
			enteredTask := scanner.Text()
			for i, v := range tasks {
				if strings.EqualFold(v, enteredTask) {
					idx = i
					break
				}
			}
			if idx == -1 {
				fmt.Printf("Your entered task '%s' is not present in tasks", enteredTask)
				return
			}
			// now remove that element with splice operator
			tasks = append(tasks[:idx], tasks[idx+1:]...)
			fmt.Println("Task Removed")
		case "3":
			fmt.Println("Enter the task, you want to mark it as complete")
			scanner.Scan()
			idx := -1
			enteredTask := scanner.Text()
			for i, v := range tasks {
				if strings.EqualFold(v, enteredTask) {
					idx = i
					break
				}
			}
			if idx == -1 {
				fmt.Printf("Your entered task '%s' is not present in tasks", enteredTask)
				return
			}
			tasks[idx] = fmt.Sprintf("%s - COMPLETED", tasks[idx])
			fmt.Println("Task Updated")
		case "4":
			if len(tasks) == 0 {
				fmt.Println("No tasks availble")
				return
			}
			fmt.Println("These are tasks for you")
			for _, task := range tasks {
				fmt.Println(task)
			}
		case "5":
			fmt.Println("Bye")
			return
		}
	}
}
