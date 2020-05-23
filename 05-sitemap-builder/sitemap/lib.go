package sitemap

import (
	"encoding/xml"
	"io"

	"github.com/Peter554/gophercises/05-sitemap-builder/links"
)

// WriteSitemap builds a sitemap from the links provided
// and writes this to a provided writer.
func WriteSitemap(ls []links.Link, w io.Writer) {
	doc := urlset{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9"}

	us := make([]url, 0)
	for _, l := range ls {
		us = append(us, url{Loc: l.Href})
	}
	doc.Urls = us

	b, e := xml.MarshalIndent(doc, "", "    ")
	if e != nil {
		panic(e)
	}

	w.Write([]byte(xml.Header))
	w.Write(b)
}

type urlset struct {
	Xmlns string `xml:"xmlns,attr"`
	Urls  []url  `xml:"url"`
}

type url struct {
	Loc string `xml:"loc"`
}
