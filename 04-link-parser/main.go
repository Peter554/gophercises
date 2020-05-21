package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	fileFlag := flag.String("file", "example.html", "The HTML file to be parsed.")
	flag.Parse()

	file, err := os.Open(*fileFlag)
	checkError(err)

	tree, err := html.Parse(file)
	checkError(err)

	links := make([]link, 0)

	addLink := func(href string, text string) {
		links = append(links, link{href: href, text: text})
	}

	visit(tree, makeVisitor(addLink))

	fmt.Println(links)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type link struct {
	href string
	text string
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
