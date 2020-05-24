package main

import "regexp"

func normalize(s string) (string, error) {
	r, e := regexp.Compile("\\D")
	if e != nil {
		return "", e
	}
	return string(r.ReplaceAll([]byte(s), []byte(""))), nil
}
