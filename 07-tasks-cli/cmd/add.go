package cmd

import (
	"strings"

	"github.com/peter554/gophercises/07-tasks-cli/db"

	"github.com/urfave/cli/v2"
)

func Add() *cli.Command {
	return &cli.Command{
		Name:   "add",
		Usage:  "Add a task to the task list",
		Action: add,
	}
}

func add(c *cli.Context) error {
	tasksService, e := db.NewTasksService()
	if e != nil {
		return e
	}
	defer tasksService.Close()

	text := strings.Join(c.Args().Slice(), " ")
	return tasksService.Add(text)
}
