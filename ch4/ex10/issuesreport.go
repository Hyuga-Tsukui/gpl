package ex10

import (
	"html/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}-------------------------------------------------------------------------
Number: {{.Number}}
User:   {{.User}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var Report = template.Must(template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))
