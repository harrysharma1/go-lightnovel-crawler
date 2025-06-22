package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Multi-URL Concurrent
type URLToBaseHTML struct {
	url  string
	html string
}

func readFile() []byte {
	data, err := os.ReadFile("lightnovel-webpages.txt")
	if err != nil {
		panic(err)
	}
	return data
}

func splitString() []string {
	data := string(readFile())
	data_split := strings.Split(data, "\n")
	return data_split
}

func checkValidHost(rawURL string) bool {
	res, err := http.Get(rawURL)
	if err != nil {
		return false
	}
	fmt.Println(res.StatusCode)

	if res.StatusCode == 401 {
		return false
	}

	return res.StatusCode == 200

}
