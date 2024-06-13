package parser

import (
    "golang.org/x/net/html"
    "io"
    "goscraper/models"
)

func ParseHTML(body io.Reader) (models.PageData, error) {
    doc, err := html.Parse(body)
    if err != nil {
        return models.PageData{}, err
    }

    titles, links := extractData(doc)
    return models.PageData{
        Titles: titles,
        Links:  links,
    }, nil
}

func extractData(n *html.Node) ([]string, []string) {
    titles := []string{}
    links := []string{}

    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
            titles = append(titles, n.FirstChild.Data)
        }
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, attr := range n.Attr {
                if attr.Key == "href" {
                    links = append(links, attr.Val)
                    break
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(n)

    return titles, links
}
