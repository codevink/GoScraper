package main

import (
    "fmt"
    "sync"
    "goscraper/fetcher"
    "goscraper/models"
)

func main() {
    urls := []string{
        "https://example.com",
        "https://golang.org",
        "https://go.dev",
    }

    var wg sync.WaitGroup
    results := make(chan models.PageData, len(urls))

    for _, url := range urls {
        wg.Add(1)
        go fetcher.FetchAndParse(url, &wg, results)
    }

    wg.Wait()
    close(results)

    for result := range results {
        fmt.Printf("URL: %s\nTitles: %v\nLinks: %v\n\n", result.URL, result.Titles, result.Links)
    }
}