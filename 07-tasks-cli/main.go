package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "tasks"
	app.Usage = "A CLI tool for managing tasks"

	app.Commands = []*cli.Command{
		{
			Name:  "list",
			Usage: "View the task list",
		},
		{
			Name:  "add",
			Usage: "Add a task to the task list",
		},
		{
			Name:  "complete",
			Usage: "Remove a task from the task list",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
