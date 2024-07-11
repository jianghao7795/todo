// Package cmd provides command-line subcommands for the gtodo application.
package cmd

import (
	"flag"
	"fmt"
	"log"
	"todo/todo"
)

func DeleteTask(todos *todo.Todos, args []string) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteID := deleteCmd.Int("id", 0, "The id of todo to be deleted")
	deleteCmd.Parse(args)

	err := todos.Delete(*deleteID)
	if err != nil {
		log.Fatal(err)
	}

	err = todos.Store(GetJsonFile())
	if err != nil {
		log.Fatal(err)
	}

	todos.Print(2, "")
	fmt.Println("Todo item deleted successfully.")
}
