package main

import ("fmt")
import ("strings")

func main() {
	fmt.Println("Please provide one or more words to search.")
}

func Parse(line string) string {
	return strings.Split(line, ";")[0]
}