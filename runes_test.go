package main

import (
	"fmt"
	"testing"
)

func Example () {
	main()
	// Oputput:
	// Please provide one or more words to search.
}

func TestParse(t *testing.T) {
	// Given
	const line= "0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;"
	// When
	gotCode, gotName := Parse(line)
	// Then
	wantCode := "0041"
	if wantCode != gotCode {
		t.Errorf("code: %v, want: %v", gotCode, wantCode)
	}
	wantName := "LATIN CAPITAL LETTER A"
	if wantName != gotName {
		t.Errorf("name: %v, want: %v", gotName, wantName)
	}
}

func TestMatch(t *testing.T) {
	testCases := []struct {
		query string
		name string
		want bool
	}{		{"BLACK", "BLACK CHESS KING", true},  
			{"WHITE", "BLACK CHESS KING", false}, 
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v in %v", tc.query, tc.name),   func(t *testing.T) {
															
																// Given
																query := tc.query
																name := tc.name
																
																// When
																got := Match(query, name)
																if got != tc.want {
																	t.Errorf("got %v; want %v", got, tc.want)
																}
															})
	}
}