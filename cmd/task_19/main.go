package main

import (
	"fmt"
	"unicode/utf8"
)

/*
Разработать программу, которая переворачивает подаваемую на вход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func main() {
	inputs := []string{
		"Hello, 世界",
		"\U0001FABF quack!",
		"this way \u261E",
	}
	for _, input := range inputs {
		fmt.Printf("\"%s\" - \"%s\"\n", input, ReverseUnicodeString(input))
	}
}

func ReverseUnicodeString(input string) string {
	unprocessed := []byte(input)
	output := make([]byte, 0, len(unprocessed))
	for len(unprocessed) > 0 {
		// decoding the last rune
		r, size := utf8.DecodeLastRune(unprocessed)
		// appending the last rune to the output
		output = utf8.AppendRune(output, r)
		// cutting off the last rune
		unprocessed = unprocessed[:len(unprocessed)-size]
	}
	return string(output)
}
