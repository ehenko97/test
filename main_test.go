package main

import (
	"testing"
)

func TestMaskLink(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http://example.com", "http://***********"},
		{"no http link here", "no http link here"},
		{"http://example.com and some text", "http://*********** and some text"},
		{"https://secure.com should not be masked", "https://secure.com should not be masked"},
		{"http://", "http://"},
		{"", ""},
	}

	for _, test := range tests {
		byteSlice := []byte(test.input)
		maskLink(byteSlice, []byte("http://"))
		result := string(byteSlice)
		if result != test.expected {
			t.Errorf("maskLink(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
