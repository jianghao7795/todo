// Package cmd provides command-line subcommands for the gtodo application.
package cmd

import (
	"flag"
	"todo/todo"
)

func ListTasks(todos *todo.Todos, args []string) {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listDone := listCmd.Int("done", 2, "The status of todo to be printed")
	listCat := listCmd.String("cat", "", "The category of tasks to be listed")

	listCmd.Parse(args)
	todos.Print(*listDone, *listCat)
}
