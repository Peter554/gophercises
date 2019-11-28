package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	fileFlag := flag.String("file", "./quiz.csv", "The path to the quiz file.")
	timeoutFlag := flag.Int("timeout", 5, "The max time to complete the quiz in secondsxs.")
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

	questions := make([]string, 0)
	answers := make([]string, 0)

	for {
		line, err := fileReader.ReadString('\n')

		if err == io.EOF {
			break
		}

		split := strings.Split(line, ",")
		questions = append(questions, split[0])
		answers = append(answers, split[1])
	}

	correct := 0
	signal := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(*timeoutFlag) * time.Second)
		signal <- "timeout"
	}()

	go func() {
		fmt.Printf("Welcome to the Quiz!\n")

		for i := 0; i < len(questions); i++ {
			fmt.Printf("Question #%d\n", i+1)
			fmt.Printf("%s\n", questions[i])
			userAnswer, _ := stdInReader.ReadString('\n')

			if answers[i] == userAnswer {
				correct++
			}
		}

		signal <- "all done"
	}()

	status := <-signal

	if status == "all done" {
		fmt.Printf("All done! You scored %d out of %d\n", correct, len(questions))
	}

	if status == "timeout" {
		fmt.Printf("Times up! You scored %d out of %d\n", correct, len(questions))
	}
}
