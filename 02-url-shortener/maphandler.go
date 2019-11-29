package main

import (
	"net/http"
)

func mapHandler(redirectMap map[string]string, fallback http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		if redirect, found := redirectMap[r.URL.Path]; found {
			http.Redirect(w, r, redirect, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}
