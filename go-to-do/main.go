package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

func askForInput(logger *log.Logger) string {
	var input string
	logger.Println("Welcome! What do you want to do? ")
	logger.Println("Options: [1] add, [2] remove")
	logger.Scanf("%s", &input)
	if input == "" {
		logger.Println(color.HiRedString("Invalid input."))
		return askForInput(logger)
	}
	return input
}

func askForTask() string {
	var input string
	fmt.Print("Enter new task here: ")

	fmt.Scanf("%s", &input)
	if input == "" {
		fmt.Println(color.HiRedString("Invalid input."))
		return askForTask()
	}
	return input
}

func main() {
	log.SetPrefix("[to-do]: ")
	logger := log.New(os.Stdout, log.Prefix(), log.Flags())
	// all tasks
	var tasks []string
	// input
	var input string = askForInput(logger)
	if input == "1" {
		var newTask string = askForTask()
		tasks = append(tasks, newTask)
		log.Print(color.HiGreenString("Task '%s' was added successfully.\n", input))
		log.Println("--- [tasks] ---")
		for i, task := range tasks { // _ is a placeholder for the index.
			log.Printf("[%d]: %s\n", i, task)
		}
	} else if input == "remove" {
		//
		return
	}

}
