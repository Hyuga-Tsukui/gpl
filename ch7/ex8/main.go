package main

import (
	"ch7-ex8/sorter"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

const (
	Title  sorter.SortKey = "Title"
	Artist sorter.SortKey = "Artist"
	Album  sorter.SortKey = "Album"
	Year   sorter.SortKey = "Year"
	Length sorter.SortKey = "Length"
)

var tracks = []*Track{
	{"Go", "Delilah", "From the Rotts Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type customSort struct {
	t      []*Track
	sorter *sorter.Sorter[Track]
}

func (c customSort) Len() int {
	return len(c.t)
}

func (c customSort) Swap(i, j int) {
	c.t[i], c.t[j] = c.t[j], c.t[i]
}

func (c customSort) Less(i, j int) bool {
	return c.sorter.Less(c.t[i], c.t[j])
}

func main() {
	sorter := sorter.NewSorter(
		[]sorter.SortKey{Title, Year},
		map[sorter.SortKey]sorter.LessFunc[Track]{
			Title: sorter.LessByField(func(t *Track) string { return t.Title }),
			Year:  sorter.LessByField(func(t *Track) int { return t.Year }),
		},
	)
	sort.Sort(customSort{
		t:      tracks,
		sorter: sorter,
	})
	printTracks(tracks)
}
