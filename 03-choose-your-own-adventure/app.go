package main

import (
	"encoding/json"
	"fmt"
	"html/template"
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
	bytes, err := ioutil.ReadFile("./story.json")

	if err != nil {
		panic(err)
	}

	pages := make(map[string]Page)
	err = json.Unmarshal(bytes, &pages)

	if err != nil {
		panic(err)
	}

	tmplt, err := template.ParseFiles("./template.html")

	if err != nil {
		panic(err)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]

		if page, found := pages[path]; found {
			tmplt.Execute(w, page)
			return
		}

		http.Redirect(w, r, "/intro", http.StatusFound)
	})

	fmt.Println("Serving on :8080")
	http.ListenAndServe(":8080", handler)
}
