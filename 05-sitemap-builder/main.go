package main

import (
	"flag"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Peter554/gophercises/05-sitemap-builder/links"
	"github.com/Peter554/gophercises/05-sitemap-builder/sitemap"
)

func main() {
	siteFlag := flag.String("site", "https://piccalil.li/", "The site for which we wish to build a sitemap")
	flag.Parse()

	links := getInternalLinks(*siteFlag)

	f, err := os.Create("./sitemap.xml")
	checkError(err)
	sitemap.WriteSitemap(links, f)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getURLBase(u string) string {
	p, e := url.Parse(u)
	checkError(e)
	return p.Scheme + "://" + p.Host
}

func isInternalLink(l links.Link, b string) bool {
	return strings.HasPrefix(l.Href, b) || strings.HasPrefix(l.Href, "/")
}

func getInternalLinks(u string) []links.Link {
	r, e := http.Get(u)
	checkError(e)
	defer r.Body.Close()
	a := links.GetLinks(r.Body)
	o := make([]links.Link, 0)
	b := getURLBase(u)
	for _, l := range a {
		if isInternalLink(l, b) {
			o = append(o, l)
		}
	}
	return o
}
