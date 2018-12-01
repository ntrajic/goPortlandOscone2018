package main

import "testing" // to kinda have a type we use in TestParse input

func Example () {
	main()
	// Oputput:
	// Please provide one or more words to search.
}

func TestParse(t *testing.T) {
	// Given
	const line= "0041;LATIN CAPITAL LETTER Ak;Lu;0;L;;;;;N;;;;0061;"
	// When
	gotCode, gotName := Parse(line)
	// Then
	wantCode := "0041"
	if wantCode != gotCode {
		t.Errorf("got: %v, want: %v", gotCode, wantCode)
	}
	wantName := "LATIN CAPITAL LETTER A"
	if wantName != gotName {
		t.Errorf("got: %v, want: %v", gotName, wantName)
	}


}