package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./quiz.csv")
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
	}
}
