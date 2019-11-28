package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fileFlag := flag.String("file", "", "The path to the quiz file.")
	helpFlag := flag.Bool("help", false, "Go for help.")
	flag.Parse()

	if *helpFlag {
		flag.PrintDefaults()
		os.Exit(0)
	}

	file, _ := os.Open(*fileFlag)
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		split := strings.Split(line, ",")
		question := split[0]
		answer := split[1]

		fmt.Println(question)
		fmt.Println(answer)
	}
}
