package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {
	var searchTerm string
	flag.StringVar(&searchTerm, "search-term", "", "The golang term to search for")
	flag.Parse()
	resp, err := http.Get("https://pkg.go.dev/compress/bzip2@latest")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(body))
}
