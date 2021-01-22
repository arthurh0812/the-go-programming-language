package github

import "time"

// IssuesURL is the URL of the Github API to search for Issues.
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult is the data structure for storing data about the resulting items when searching with the IssuesURL.
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue is the data structure for storing the data of a single Issue.
type Issue struct {
	Number    int    `json:"number"`
	HTMLURL   string `json:"html_url"`
	Title     string `json:"title"`
	State     string
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in markdown format
}

// User is the data structure for storing the data of a single user.
type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}
