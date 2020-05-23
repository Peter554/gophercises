package links

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link in a web page.
type Link struct {
	Href string
	Text string
}

// GetLinks gets all the links in the provided HTML document.
func GetLinks(r io.Reader) []Link {
	tree, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	links := make([]Link, 0)

	addLink := func(href string, text string) {
		links = append(links, Link{Href: href, Text: text})
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

func makeVisitor(addLink func(string, string)) func(*html.Node) {
	visitor := func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			href := getHref(node)
			text := ""
			getText(node, &text)
			addLink(href, text)
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

func getText(node *html.Node, text *string) {
	if node.Type == html.TextNode {
		s := strings.TrimSpace(node.Data)
		if len(*text) > 0 {
			*text += " "
		}
		*text += s
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		getText(child, text)
	}
}
