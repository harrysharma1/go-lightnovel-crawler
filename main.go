package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

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

func htmlRender(i int, urlStruct chan URLToBaseHTML) {
	data := splitString()
	url := data[i]

	client := http.Client{}
	req, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req.Request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	resp_body, _ := io.ReadAll(resp.Body)
	urlStruct <- URLToBaseHTML{url: url, html: string(resp_body)}
}

func main() {
	data := splitString()
	urlStruct := make(chan URLToBaseHTML)
	for i := 0; i < len(data); i++ {
		go htmlRender(i, urlStruct)
	}

	result := []URLToBaseHTML{}
	for i := 0; i < len(data); i++ {
		fmt.Printf("URL %d", i, "done`")
		res := <-urlStruct
		result = append(result, res)

	}

	close(urlStruct)
	for _, j := range result {
		fmt.Println(j, "\n")
	}
}
