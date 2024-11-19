package main

import (
	"fmt"
	"os"

	"awesomeProject/service"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file> [<output_file>]")
		return
	}

	inputFile := os.Args[1]
	outputFile := "output.txt"
	if len(os.Args) >= 3 {
		outputFile = os.Args[2]
	}

	prod := service.NewFileProducer(inputFile)
	pres := service.NewFilePresenter(outputFile)

	srv := service.NewService(prod, pres)
	if err := srv.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func maskLink(byteSlice, mask []byte) {
	httpLength := len(mask)
	for i := 0; i <= len(byteSlice)-httpLength; i++ {
		match := true
		for j := 0; j < httpLength; j++ {
			if byteSlice[i+j] != mask[j] {
				match = false
				break
			}
		}
		if match {
			endIndex := i + httpLength
			for endIndex < len(byteSlice) && (byteSlice[endIndex] >= 'a' && byteSlice[endIndex] <= 'z' ||
				byteSlice[endIndex] >= 'A' && byteSlice[endIndex] <= 'Z' ||
				byteSlice[endIndex] >= '0' && byteSlice[endIndex] <= '9' ||
				byteSlice[endIndex] == '.' || byteSlice[endIndex] == '/' || byteSlice[endIndex] == ':' || byteSlice[endIndex] == '?') {
				endIndex++
			}
			for k := i + httpLength; k < endIndex; k++ {
				byteSlice[k] = '*'
			}
		}
	}
}
