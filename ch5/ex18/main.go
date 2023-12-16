package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	url := "https://example.com"

	filename, n, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("filenam: %q size: %q", filename, n)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	n, err = io.Copy(f, resp.Body)
	return local, n, err
}
