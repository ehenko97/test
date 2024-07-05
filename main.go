package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Чтение строки из стандартного ввода
	fmt.Print("INPUT: ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// Преобразование строки в срез байтов
	byteSlice := []byte(text)

	// Маскирование ссылок, начинающихся с "http://"
	maskLink(byteSlice, []byte("http://"))

	// Печать измененного байтового среза как строки
	fmt.Printf("OUTPUT: %s\n", string(byteSlice))
}

// Функция для маскирования ссылок, начинающихся с "http://"
func maskLink(byteSlice, mask []byte) {
	// Длина подстроки "http://"
	httpLength := len(mask)

	// Проход по байтовому срезу и поиск подстрок, начинающихся с "http://"
	for i := 0; i <= len(byteSlice)-httpLength; i++ {
		match := true

		// Проверка совпадения с "http://"
		for j := 0; j < httpLength; j++ {
			if byteSlice[i+j] != mask[j] {
				match = false
				break
			}
		}
		// Если найдено совпадение, замена последующих символов на '*'
		if match {
			// Индекс конца ссылки
			endIndex := i + httpLength
			for endIndex < len(byteSlice) && (byteSlice[endIndex] >= 'a' && byteSlice[endIndex] <= 'z' ||
				byteSlice[endIndex] >= 'A' && byteSlice[endIndex] <= 'Z' ||
				byteSlice[endIndex] >= '0' && byteSlice[endIndex] <= '9' ||
				byteSlice[endIndex] == '.' || byteSlice[endIndex] == '/' || byteSlice[endIndex] == ':' || byteSlice[endIndex] == '?') {
				endIndex++
			}

			// Маскирование символов
			for k := i + httpLength; k < endIndex; k++ {
				byteSlice[k] = '*'
			}
			// Прерывание после первого совпадения
			break
		}
	}
}
