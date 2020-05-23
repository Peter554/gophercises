package links

import (
	"io"

	"golang.org/x/net/html"
)

// Link represents a link in a web page.
type Link struct {
	Href string
}

// GetLinks gets all the links in the provided HTML document.
func GetLinks(r io.Reader) []Link {
	tree, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	links := make([]Link, 0)

	addLink := func(href string) {
		links = append(links, Link{Href: href})
	}

	visit(tree, makeVisitor(addLink))

	return links
}

func visit(tree *html.Node, visitor func(*html.Node)) {
	visitor(tree)
	for child := tree.FirstChild; child != nil; child = child.NextSibling {
		visit(child, visitor)
	}
}

func makeVisitor(addLink func(string)) func(*html.Node) {
	visitor := func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			href := getHref(node)
			addLink(href)
		}
	}
	return visitor
}

func getHref(node *html.Node) string {
	attributes := node.Attr
	for _, attribute := range attributes {
		if attribute.Key == "href" {
			return attribute.Val
		}
	}
	return ""
}
