package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues?q=go"

type IssuesSearchResult struct {
	TotalCount        int  `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items             []Issue
}

type Issue struct {
	ID        int
	NodeID    string `json:"node_id"`
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	CreatedAt time.Time `json:"created_at"`
}

var result IssuesSearchResult

func main() {
	//请求接口获取json,并解析
	resp, err := http.Get(IssuesURL)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Printf("search query failed: %s", resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		fmt.Println(err)
	}
	resp.Body.Close()
	fmt.Println(result.Items[0].NodeID)
}
