package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func main() {
	fileFlag := flag.String("csv", "./quiz.csv", "The path to the quiz CSV.")
	timeoutFlag := flag.Int("timeout", 30, "The max time to complete the quiz in seconds.")
	flag.Parse()

	file, _ := os.Open(*fileFlag)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, _ := csvReader.ReadAll()

	appState := make(chan string, 1)
	correct := 0

	go func() {
		time.Sleep(time.Duration(*timeoutFlag) * time.Second)
		appState <- "timeout"
	}()

	go func() {
		stdInReader := bufio.NewReader(os.Stdin)
		fmt.Printf("Welcome to the Quiz!\n")

		for idx, record := range records {
			fmt.Printf("Question #%d\n", idx+1)
			fmt.Printf("%s\n", record[0])
			userAnswer, _ := stdInReader.ReadString('\n')

			if normalize(record[1]) == normalize(userAnswer) {
				fmt.Printf("Correct!\n")
				correct++
			} else {
				fmt.Printf("Ooops!\n")
			}
		}

		appState <- "all done"
	}()

	state := <-appState

	if state == "all done" {
		fmt.Printf("All done! You scored %d out of %d\n", correct, len(records))
	}

	if state == "timeout" {
		fmt.Printf("Times up! You scored %d out of %d\n", correct, len(records))
	}
}
