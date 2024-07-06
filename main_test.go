package main

import (
	"io/ioutil"
	"testing"
)

// Структура для хранения тестовых данных
type testData struct {
	fileName string
	expected string
}

func TestMaskLink(t *testing.T) {
	tests := map[string]testData{
		"Test 1": {"testdata/test1.txt", "Check this http://****************"},
		"Test 2": {"testdata/test2.txt", "No link here"},
		"Test 3": {"testdata/test3.txt", "Multiple links http://*********** and http://********"},
		"Test 4": {"testdata/test4.txt", "Edge case http://"},
		"Test 5": {"testdata/test5.txt", "Trailing http://*********** and some text"},
	}

	for name, test := range tests {
		// Чтение содержимого файла тестовых данных
		content, err := ioutil.ReadFile(test.fileName)
		if err != nil {
			t.Fatalf("Failed to read test data file '%s': %s", test.fileName, err)
		}

		byteSlice := content
		maskLink(byteSlice, []byte("http://"))
		output := string(byteSlice)

		if output != test.expected {
			t.Errorf("[%s] Expected '%s', but got '%s'", name, test.expected, output)
		}
	}
}
