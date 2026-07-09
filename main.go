package main

import (
	"ABC-Live-Web-API/fetchers"
	"fmt"
)

func main() {
	data := fetchers.ETFetcher()
	fmt.Println(data)
}