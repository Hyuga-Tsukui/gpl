package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {

	var w WordCounter

	fmt.Fprintf(&w, "hello world! hoge hoge %s", "fuga fuga")

	fmt.Println(w)

    fmt.Println("-------7.1")

    cw, c := CountingWriter(&w)

    fmt.Fprintf(cw, "piyo")
    fmt.Println(*c)
    fmt.Println(w)
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var b int64
    
    return &WrapWriter{
        Writer: w,
        Counter: &b,
    }, &b
}

type WrapWriter struct {
    Writer io.Writer
    Counter *int64
}

func (w *WrapWriter) Write(p []byte) (int, error) {
    *w.Counter += int64(len(p))
    return w.Writer.Write(p)
}
