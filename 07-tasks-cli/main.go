package main

import (
	"log"
	"os"

	"github.com/peter554/gophercises/07-tasks-cli/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "tasks"
	app.Usage = "A CLI tool for managing tasks"

	app.Commands = []*cli.Command{
		cmd.List(),
		cmd.Add(),
		cmd.Complete(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
