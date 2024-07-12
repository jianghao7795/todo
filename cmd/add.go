// Package cmd provides command-line subcommands for the gtodo application.
package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jianghao7795/todo/todo"
)

func AddTask(todos *todo.Todos, args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTask := addCmd.String("task", "", "The content of new todo item")
	addCat := addCmd.String("cat", "Uncategorized", "The category of the todo item")

	addCmd.Parse(args)

	if len(*addTask) == 0 {
		fmt.Println("Error: the --task flag is required for the 'add' subcommand.")
		os.Exit(1)
	}
	todos.Add(*addTask, *addCat)
	err := todos.Store(GetJsonFile())
	if err != nil {
		log.Fatal(err)
	}

	todos.Print(2, "")
	fmt.Println("Todo item added successfully.")
}
