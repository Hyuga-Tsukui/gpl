package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

type clock struct {
	loc  string
	port string
}

type clocks []clock

func (i *clocks) Set(v string) error {
	s := strings.Split(v, "=")
	*i = append(*i, clock{s[0], s[1]})
	return nil
}

func (i *clocks) String() string {
	return fmt.Sprintf("%v", *i)
}

func flagClocks(name string, usage string) *clocks {
	c := new(clocks)
	flag.Var(c, name, usage)
	return c
}

var clocksFlag = flagClocks("clocks", "clocks to run")

func main() {
	flag.Parse()

	if len(*clocksFlag) == 0 {
		log.Fatal("no clocks provided")
	}

	strCh := make(chan string, 10)
	defer close(strCh)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go spinner(ctx)

	for _, c := range *clocksFlag {
		go muxClocks(c, strCh)
	}

	// 表示領域をクリア.
	fmt.Print("\033[2J")

	clockMap := make(map[string]string)
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case str := <-strCh:
			parts := strings.SplitN(str, ": ", 2)
			if len(parts) == 2 {
				clockMap[parts[0]] = parts[1]
			}
		case <-ticker.C:
			cancel()
			// 左上から入力するためカーソルを先頭に移動.
			fmt.Print("\033[H")

			for _, c := range *clocksFlag {
				fmt.Printf("%s: %s", c.loc, clockMap[c.loc])
			}
		}
	}
}

func spinner(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			for _, r := range `-\|/` {
				fmt.Print("\033[H")
				fmt.Printf("\r%c", r)

				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

func muxClocks(clock clock, ch chan string) {
	conn, err := net.Dial("tcp", "localhost"+clock.port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Fatal("connection closed")
			}
			break
		}
		ch <- clock.loc + ": " + string(buf[:n])
	}
}

func mustCopy(ch chan string, src io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := src.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Fatal("connection closed")
			}
			break
		}
		ch <- string(buf[:n])
	}
}
