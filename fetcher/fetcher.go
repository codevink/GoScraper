package fetcher

import (
    "fmt"
    "net/http"
    "sync"
    "goscraper/models"
    "goscraper/parser"
)

func FetchAndParse(url string, wg *sync.WaitGroup, ch chan<- models.PageData) {
    defer wg.Done()

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error fetching URL:", err)
        return
    }
    defer resp.Body.Close()

    pageData, err := parser.ParseHTML(resp.Body)
    if err != nil {
        fmt.Println("Error parsing HTML:", err)
        return
    }

    pageData.URL = url
    ch <- pageData
}
