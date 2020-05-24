package cmd

import (
	"fmt"

	"github.com/peter554/gophercises/07-tasks-cli/db"

	"github.com/urfave/cli/v2"
)

func List() *cli.Command {
	return &cli.Command{
		Name:   "list",
		Usage:  "View the task list",
		Action: list,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "Includes completed tasks in the list",
			},
		},
	}
}

func list(c *cli.Context) error {
	tasksService, e := db.NewTasksService()
	if e != nil {
		return e
	}
	defer tasksService.Close()

	tasks, e := tasksService.GetAll()
	if e != nil {
		return e
	}

	all := c.Bool("all")
	for _, t := range tasks {
		if all {
			if t.Completed {
				fmt.Printf("(%d) X %s\n", t.ID, t.Text)
			} else {
				fmt.Printf("(%d)   %s\n", t.ID, t.Text)
			}
		} else if !t.Completed {
			fmt.Printf("(%d) %s\n", t.ID, t.Text)
		}
	}
	return nil
}
