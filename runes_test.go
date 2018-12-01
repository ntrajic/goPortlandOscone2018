package main

func Example () {
	main()
	// Oputput:
	// Please provide one or more words to search.
}

func TestParse(t *testing.T) {
	// Given
	const line= "0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;"
	// When
	got := Parse(line)
	// Then
	want := "0041"
	if want != got {
		t.Errorf("got: %v, want: %v", got, want)
	}


}