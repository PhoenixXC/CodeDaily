package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"
)

const (
	IssuesURL = "https://api.github.com/search/issues"
	templ     = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`
)

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"` // 通过 time.Time.MarshalJSON 来实现自定义序列化接口
	Body      string
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

var report = template.Must(template.New("issueList").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%d issues:\n", result.TotalCount)
	// for _, item := range result.Items {
	//     // Width is specified by an optional decimal number immediately preceding the verb.
	//     // If absent, the width is whatever is necessary to represent the value.
	//     // ....
	//     // For strings, byte slices and byte arrays, however,
	//     // precision limits the length of the input to be formatted
	//     // (not the size of the output), truncating if necessary.
	//     // 对于字符串而言，%9.12s 限制宽度最少为9，不够空格补全，宽度最多为12，太多了就截断
	//     fmt.Printf("#%-5d, %50.50s, %10.55s\n", item.Number, item.User.HTMLURL, item.Title)
	// }
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		_ = resp.Body.Close()
		return nil, err
	}
	_ = resp.Body.Close()
	return &result, nil
}
