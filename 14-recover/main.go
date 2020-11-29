package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/panic/", func(w http.ResponseWriter, r *http.Request) {
		panic("Oh no!")
	})

	mux.HandleFunc("/panic-after/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hello!</h1>")
		panic("Oh no!")
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hello!</h1>")
	})

	log.Fatal(http.ListenAndServe(":3000", withRecovery(mux, true)))
}

func withRecovery(handler http.Handler, dev bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				stack := string(debug.Stack())

				log.Println(err)
				log.Println(stack)

				w.WriteHeader(http.StatusInternalServerError)
				if dev {
					fmt.Fprintf(w, "<p>%v</p><p>%v<p/>", err, stack)
				} else {
					fmt.Fprint(w, "<p>Oops, something went wrong! :(</p>")
				}
			}
		}()

		nw := &delayedResponseWriter{ResponseWriter: w}
		handler.ServeHTTP(nw, r)
		err := nw.write()
		if err != nil {
			panic(err)
		}
	})
}

type delayedResponseWriter struct {
	http.ResponseWriter
	writes [][]byte
	status int
}

func (rw *delayedResponseWriter) Write(b []byte) (int, error) {
	rw.writes = append(rw.writes, b)
	return len(b), nil
}

func (rw *delayedResponseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
}

func (rw *delayedResponseWriter) write() error {
	if rw.status != 0 {
		rw.ResponseWriter.WriteHeader(rw.status)
	}
	for _, b := range rw.writes {
		_, err := rw.ResponseWriter.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}
