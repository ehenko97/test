package main

import (
	"testing"
)

func TestMaskLink(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Mask1", "http://example.com", "http://***********"},
		{"Mask2", "no http link here", "no http link here"},
		{"Mask3", "http://example.com and some text", "http://*********** and some text"},
		{"Mask4", "https://secure.com should not be masked", "https://secure.com should not be masked"},
		{"Mask5", "http://", "http://"},
		{"Mask6", "", ""},
	}

	for _, test := range tests {
		byteSlice := []byte(test.input)
		maskLink(byteSlice, []byte("http://"))
		result := string(byteSlice)
		if result != test.expected {
			t.Errorf("Test case %q: maskLink(%q) = %q; want %q", test.name, test.input, result, test.expected)
		}
	}
}
