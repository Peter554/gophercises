package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func main() {
	(&cli.App{
		Name:  "rename",
		Usage: "Tool to rename files",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "dir",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			dir := c.String("dir")

			dir, err := filepath.Abs(dir)
			if err != nil {
				log.Fatal(err)
				return err
			}

			fi, err := os.Stat(dir)
			if err != nil {
				log.Fatal(err)
				return err
			}
			if !fi.IsDir() {
				err := fmt.Errorf("%s is not a directory", dir)
				log.Fatal(err)
				return err
			}

			counts := map[string]int{}

			filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				parent := filepath.Dir(path)
				if _, found := counts[parent]; found {
					counts[parent]++
				} else {
					counts[parent] = 1
				}
				count, _ := counts[parent]
				newPath := filepath.Join(parent, fmt.Sprintf("%s_%03d%s", filepath.Base(parent), count, filepath.Ext(path)))
				return os.Rename(path, newPath)
			})
			return nil
		},
	}).Run(os.Args)
}
