package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type crawledData struct {
	url             string
	textBasedHTMl   string
	parserBasedHTML *html.Node
}

func listAllLinks(rawURL string) []string {
	url, found := strings.CutSuffix(rawURL, "/")
	if !found {
		fmt.Println("Error: removing leading /")
	}

	res, err := http.Get(rawURL)
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		panic(err)
	}

	return traverse_nodes(url, doc)

}

func traverse_nodes(urlPrefix string, node *html.Node) []string {
	var links []string
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				if a.Val != "javascript:void(0)" {
					links = append(links, urlPrefix+a.Val)
					break
				}
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		childLinks := traverse_nodes(urlPrefix, c)
		links = append(links, childLinks...)
	}
	return links
}

func store() {
	const file = "crawledData.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic(err)
	}
	fmt.Print(db)

}

func crawler(rawURL string) {
	hrefs := listAllLinks(rawURL)
	var data = []crawledData{}
	for i := range hrefs {
		res, err := http.Get(hrefs[i])
		if err != nil {
			panic(err)
		}
		htmlBody, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		doc, err := html.Parse(res.Body)
		if err != nil {
			panic(err)
		}

		data = append(data, crawledData{url: hrefs[i], textBasedHTMl: string(htmlBody), parserBasedHTML: doc})
	}
	store()
}
