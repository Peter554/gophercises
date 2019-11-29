package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := defaultHandler()

	m := map[string]string{
		"/peter554": "https://github.com/peter554",
	}

	handler = MapHandler(m, handler)

	fmt.Println("Serving on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Go!")
}
