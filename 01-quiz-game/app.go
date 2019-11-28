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

	fileReader := bufio.NewReader(file)
	stdInReader := bufio.NewReader(os.Stdin)

	completed := 0
	correct := 0

	fmt.Printf("Welcome to the Quiz!\n")

	for {
		line, err := fileReader.ReadString('\n')

		if err == io.EOF {
			break
		}

		split := strings.Split(line, ",")
		question := split[0]
		answer := split[1]

		fmt.Printf("Question #%d\n", completed+1)
		fmt.Printf("%s\n", question)
		userAnswer, _ := stdInReader.ReadString('\n')

		completed++

		if answer == userAnswer {
			correct++
		}
	}

	fmt.Printf("All done! You scored %d out of %d\n", correct, completed)
}
