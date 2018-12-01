package main

import (
	"fmt"
	"strings"
)


func main() {
	fmt.Println("Please provide one or more words to search.")
}

func Parse(line string) (string, string) {
	a := strings.Split(line, ";")
	return a[0], a[1]
}

func Match(query string, name string) bool {
	queryTerms := strings.Split(strings.ToUpper(query),"")
	for _, queryTerm := range queryTerms {
		if !strings.Contains(name,queryTerm) {
			return false
		}
	}
	return true
}