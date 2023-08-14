package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
	abcd — true
	abCdefAaf — false
	aabcd — false
*/

func stringRunesAreUnique(input string) bool {
	// a set of seen runes
	seen := make(map[rune]struct{})
	// converting to bytes to simplify working with runes
	unprocessed := []byte(input)
	for len(unprocessed) > 0 {
		r, size := utf8.DecodeRune(unprocessed)
		// using unicode.ToLower to handle unicode symbols
		r = unicode.ToLower(r)
		// return false if already seen, add to seen otherwise
		if _, found := seen[r]; found {
			return false
		}
		seen[r] = struct{}{}
		// skipping the processed part (one rune)
		unprocessed = unprocessed[size:]
	}
	return true
}

func main() {
	inputs := []string{
		"abcd", "abCdefAaf", "aabcd", "утка", "Уутка",
	}
	for _, input := range inputs {
		fmt.Println(input, stringRunesAreUnique(input))
	}
}
