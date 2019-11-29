package main

import (
	"net/http"
)

func MapHandler(m map[string]string, fallback http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if redirect, found := m[r.URL.Path]; found {
			http.Redirect(w, r, redirect, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	})

}
