package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func crawler(rawURL string) {
	res, err := http.Get(rawURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	s := doc.Data
	fmt.Println(s)

}
