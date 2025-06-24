package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func crawler(rawURL string) {
	url, found := strings.CutSuffix(rawURL, "/")
	if !found {
		fmt.Println("Error: removing leading /")
	}
	fmt.Print(url)

	res, err := http.Get(rawURL)
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		panic(err)
	}
	traverse_nodes(url, doc)
}

func traverse_nodes(urlPrefix string, node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				fmt.Printf("href: %s%s\n", urlPrefix, a.Val)
				break
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		traverse_nodes(urlPrefix, c)
	}
}
