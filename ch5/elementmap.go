package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Chapter5 5.2
func MapElementCount(source string) (map[string]int, error) {
	result := map[string]int{}
	doc, err := html.Parse(NewReader(source))
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

// Chapter7 7.4
type Reader struct {
	s string
	i int64
}

func (r *Reader) Read(b []byte) (int, error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}

	n := copy(b, r.s[r.i:])
	r.i += int64(n)

	return n, nil
}

func NewReader(s string) *Reader {
	return &Reader{s: s, i: 0}
}
