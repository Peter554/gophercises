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
	site := flag.String("site", "https://piccalil.li/", "The site for which we wish to build a sitemap")
	flag.Parse()

	ls := links.NewLinkSet()
	collectSiteLinks(*site, ls)

	f, err := os.Create("./sitemap.xml")
	checkError(err)
	sitemap.WriteSitemap(ls.Values(), f)
}

func collectSiteLinks(site string, ls links.LinkSet) {
	for _, l := range getInternalLinks(site) {
		if ls.Contains(l) {
			continue
		}
		ls.Add(l)
		collectSiteLinks(l.Href, ls)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getInternalLinks(u string) []links.Link {
	r, e := http.Get(u)
	checkError(e)
	defer r.Body.Close()
	a := links.GetLinks(r.Body)
	o := make([]links.Link, 0)
	b := getURLBase(u)
	for _, l := range a {
		if strings.HasPrefix(l.Href, "/") {
			l.Href = b + l.Href
		}
		if strings.HasPrefix(l.Href, b) {
			o = append(o, l)
		}
	}
	return o
}

func getURLBase(u string) string {
	p, e := url.Parse(u)
	checkError(e)
	o := &url.URL{
		Scheme: p.Scheme,
		Host:   p.Host,
	}
	return o.String()
}
