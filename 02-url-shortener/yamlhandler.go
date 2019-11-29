package main

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

func yamlHandler(redirectYAML string, fallback http.Handler) http.Handler {
	redirects := parseYAML(redirectYAML)
	redirectMap := toMap(redirects)

	return mapHandler(redirectMap, fallback)
}

type redirect struct {
	Path        string `yaml:"path"`
	RedirectURL string `yaml:"redirectURL"`
}

func parseYAML(s string) []redirect {
	var out []redirect

	err := yaml.Unmarshal([]byte(s), &out)

	if err != nil {
		panic(err)
	}

	return out
}

func toMap(redirects []redirect) map[string]string {
	out := make(map[string]string)

	for _, r := range redirects {
		out[r.Path] = r.RedirectURL
	}

	return out
}
