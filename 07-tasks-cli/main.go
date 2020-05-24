package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
	bolt "go.etcd.io/bbolt"
)

func main() {
	app := cli.NewApp()
	app.Name = "tasks"
	app.Usage = "A CLI tool for managing tasks"

	app.Commands = []*cli.Command{
		{
			Name:   "list",
			Usage:  "View the task list",
			Action: listCmd,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "all",
					Value: false,
					Usage: "Includes completed tasks in the list",
				},
			},
		},
		{
			Name:   "add",
			Usage:  "Add a task to the task list",
			Action: addCmd,
		},
		{
			Name:   "complete",
			Usage:  "Mark a task from the task list as completed",
			Action: completeCmd,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type task struct {
	ID        uint64
	Text      string
	Completed bool
}

func listCmd(c *cli.Context) error {
	all := c.Bool("all")

	db, e := bolt.Open("tasks.db", 0666, nil)
	if e != nil {
		return e
	}
	defer db.Close()

	tasks := make([]task, 0)
	e = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		if b == nil {
			return errors.New("bucket does not exist, you probably didn't add any tasks yet")
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			t := &task{}
			e := json.Unmarshal(v, t)
			if e != nil {
				return e
			}
			tasks = append(tasks, *t)
		}
		return nil
	})
	if e != nil {
		return e
	}

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

func addCmd(c *cli.Context) error {
	db, e := bolt.Open("tasks.db", 0666, nil)
	if e != nil {
		return e
	}
	defer db.Close()

	text := strings.Join(c.Args().Slice(), " ")
	t := task{Text: text, Completed: false}

	db.Update(func(tx *bolt.Tx) error {
		b, e := tx.CreateBucketIfNotExists([]byte("tasks"))
		if e != nil {
			return e
		}
		id, e := b.NextSequence()
		t.ID = id
		if e != nil {
			return e
		}
		j, e := json.Marshal(t)
		if e != nil {
			return e
		}
		return b.Put(itob(id), j)
	})

	return nil
}

func completeCmd(c *cli.Context) error {
	db, e := bolt.Open("tasks.db", 0666, nil)
	if e != nil {
		return e
	}
	defer db.Close()

	id, e := strconv.Atoi(c.Args().First())
	if e != nil {
		return e
	}

	return db.Update(func(tx *bolt.Tx) error {
		b, e := tx.CreateBucketIfNotExists([]byte("tasks"))
		if e != nil {
			return e
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if btoi(k) == uint64(id) {
				t := &task{}
				json.Unmarshal(v, t)
				t.Completed = true
				s, e := json.Marshal(t)
				if e != nil {
					return e
				}
				return b.Put(k, s)
			}
		}
		return nil
	})
}

func itob(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

func btoi(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}
