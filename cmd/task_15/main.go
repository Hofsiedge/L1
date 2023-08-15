package main

import (
	"errors"
	"strings"
	"unicode/utf8"
)

/*
К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

// I guess the argument to createHugeString is supposed to be length
func createHugeString(length int) string {
	return strings.Repeat(".", length)
}

var justString string

func someFunc() {
	// if 1 << 10 is string length, then 100 would be enough
	v := createHugeString(1 << 10)
	// if createHugeString accepts resulting length
	// this is an unnecessary copy - could just assign it right away
	justString = v[:100]
}

// perhaps, something like this would be a better version
func correctedFunc() {
	justString = createHugeString(100)
}

// correctedWithUnicode checks input length and works with unicode
func correctedWithUnicode() (string, error) {
	// if arg is not just resulting length
	v := createHugeString(1 << 10)
	// if v contains non-ASCII symbols can't just take a subslice of v
	input := []byte(v)
	output := make([]byte, 0, 100*4) // enough for 100 32-bit elements
	for i := 0; i < 100; i++ {
		if len(input) == 0 {
			return "", errors.New("not enough symbols")
		}
		// take the first rune
		r, size := utf8.DecodeRune(input)
		// append it to the output
		utf8.AppendRune(output, r)
		// reslice input to skip the rune
		input = input[size:]
	}
	return string(output), nil
}

func main() {
	someFunc()
}
