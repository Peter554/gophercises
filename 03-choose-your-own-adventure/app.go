package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Page struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Link string `json:"arc"`
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

	handler := story_handler(pages, tmplt)
	handler = static_handler(handler)

	fmt.Println("Serving on :8080")
	http.ListenAndServe(":8080", handler)
}

func story_handler(pages map[string]Page, tmplt *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]

		if page, found := pages[path]; found {
			tmplt.Execute(w, page)
			return
		}

		http.Redirect(w, r, "/intro", http.StatusFound)
	})
}

func static_handler(fallback http.Handler) http.Handler {
	fs := http.FileServer(http.Dir("./static"))
	fs = http.StripPrefix("/static", fs)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if matched, _ := regexp.Match("/static", []byte(r.URL.Path)); matched {
			fs.ServeHTTP(w, r)
			return
		}

		fallback.ServeHTTP(w, r)
	})
}
