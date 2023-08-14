package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	input := "snow dog sun 世 界"
	fmt.Println(ReverseWordsWithStrings(input))
}

func ReverseWordsWithStrings(input string) string {
	words := strings.Fields(input)
	L := len(words)
	// reverse the slice
	for i := 0; i < L/2; i++ {
		// swapping words
		words[i], words[L-i-1] = words[L-i-1], words[i]
	}
	return strings.Join(words, " ")
}
