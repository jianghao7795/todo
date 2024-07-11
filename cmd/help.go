// Package cmd provides command-line subcommands for the gtodo application.
package cmd

import (
	"fmt"
)

func Help() {
	fmt.Println("Welcome to gtodo CLI app!")
	fmt.Println("Usage: gtodo <command> [arguments]")
	fmt.Println("Available commands:")
	fmt.Println("  init                		Create an empty JSON file to store tasks")
	fmt.Println("  add <task> <cat>    		Add a new task")
	fmt.Println("  list <done>         		List all tasks")
	fmt.Println("  update <id> <task> <cat>	Update an existing task")
	fmt.Println("  delete <id>         		Delete an existing task")
	fmt.Println("")
}
