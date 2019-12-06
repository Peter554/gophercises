package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title   string   `json"title"`
	Story   []string `json"story"`
	Options []Option `json"options"`
}

type Option struct {
	Text string `json"text"`
	Arc  string `json"arc"`
}

func main() {
	bytes, _ := ioutil.ReadFile("./story.json")

	data := make(map[string]Page)
	err := json.Unmarshal(bytes, &data)

	if err != nil {
		panic(err)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(data["intro"].Title))
	})

	fmt.Println("Serving on :8080")
	http.ListenAndServe(":8080", handler)
}
