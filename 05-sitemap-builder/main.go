package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Peter554/gophercises/05-sitemap-builder/links"
)

func main() {
	siteFlag := flag.String("site", "https://piccalil.li/", "The site for which we wish to build a sitemap")
	flag.Parse()

	links := getLinks(*siteFlag)

	fmt.Println(links)
}

func getLinks(url string) []links.Link {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return links.GetLinks(resp.Body)
}
