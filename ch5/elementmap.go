package main

import (
	"fmt"
    "io"

	"golang.org/x/net/html"
)

// Chapter5 5.2
func MapElementCount(source io.Reader) (map[string]int, error) {
	result := map[string]int{}
	doc, err := html.Parse(source)
	if err != nil {
		return result, fmt.Errorf("failed to parse source by %s", err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			result[n.Data]++
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return result, nil
}
