package cmd

import (
	"strconv"

	"github.com/peter554/gophercises/07-tasks-cli/db"
	"github.com/urfave/cli/v2"
)

func Complete() *cli.Command {
	return &cli.Command{
		Name:   "complete",
		Usage:  "Mark a task from the task list as completed",
		Action: complete,
	}
}

func complete(c *cli.Context) error {
	tasksService, e := db.NewTasksService()
	if e != nil {
		return e
	}
	defer tasksService.Close()

	id, e := strconv.Atoi(c.Args().First())
	if e != nil {
		return e
	}

	return tasksService.Complete(uint64(id))
}
