package link

import (
    "golang.org/x/net/html"
    "io"
    "strings"
)

// Link represents an extracted link
type Link struct {
    Href string
    Text string
}

// Extract takes an HTML document and returns a slice of links
func Extract(r io.Reader) ([]Link, error) {
    doc, err := html.Parse(r)
    if err != nil {
        return nil, err
    }
    var links []Link
    visitNode := func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            link := Link{}
            for _, attr := range n.Attr {
                if attr.Key == "href" {
                    link.Href = attr.Val
                    break
                }
            }
            link.Text = getText(n)
            links = append(links, link)
        }
    }
    forEachNode(doc, visitNode, nil)
    return links, nil
}

func getText(n *html.Node) string {
    if n.Type == html.TextNode {
        return n.Data
    }
    if n.Type != html.ElementNode {
        return ""
    }
    var text string
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        text += getText(c) + " "
    }
    return strings.Join(strings.Fields(text), " ")
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
    if pre != nil {
        pre(n)
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre, post)
    }
    if post != nil {
        post(n)
    }
}
