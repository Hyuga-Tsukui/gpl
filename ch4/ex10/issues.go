package ex10

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssueURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssueSearchResult, *IssueSearchResult, *IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q + "&per_page=100&page=1")
	if err != nil {
		return nil, nil, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, nil, nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, nil, nil, err
	}
	resp.Body.Close()

	var (
		short  []Issue
		medium []Issue
		long   []Issue
	)
	baseTime := time.Now()
	monthAgo := baseTime.AddDate(0, -1, 0)
	yearAgo := baseTime.AddDate(-1, 0, 0)
	for _, i := range result.Items {
		// 1ヶ月未満.
		if i.CreatedAt.Compare(monthAgo) > 0 {
			short = append(short, i)
			continue
		}
		// 1年未満.
		if i.CreatedAt.Compare(yearAgo) > 0 {
			medium = append(medium, i)
			continue
		}
		long = append(long, i)
	}
	shortResult := &IssueSearchResult{
		TotalCount: len(short),
		Items:      short,
	}
	mediumResult := &IssueSearchResult{
		TotalCount: len(medium),
		Items:      medium,
	}
	longResult := &IssueSearchResult{
		TotalCount: len(long),
		Items:      long,
	}
	fmt.Printf("----- Less than one month -----\n")
	printItems(short)
	fmt.Printf("----- Less than one year -----\n")
	printItems(medium)
	fmt.Printf("----- More than one year -----\n")
	printItems(long)
	return shortResult, mediumResult, longResult, nil
}

func printItems(issues []Issue) {
	for _, i := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n", i.Number, i.User.Login, i.Title)
	}
}
