package main

import (
	"strconv"
	"bufio"
	"io"
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
	if query == "" {
		return false
	}
	queryTerms := strings.Split(strings.ToUpper(query),"")
	for _, queryTerm := range queryTerms {
		if !strings.Contains(name,queryTerm) {
			return false
		}
	}
	return true
}

type CodeName struct {
	Code string
	Name string
}

func Select (data io.Reader, query string) []CodeName {
	scanner := bufio.NewScanner(data)
	result := []CodeName{}
	for scanner.Scan() {
		code, name := Parse(scanner.Text())
		if Match(query, name) {
			result = append(result, CodeName{code, name})
		}
	}
	return result
}

func StringToRune(strHex string) rune {
	code, err := strconv.ParseInt(strHex, 16, 32)
	if err != nil {	
		panic(err)
	}
	return rune(code)
}
