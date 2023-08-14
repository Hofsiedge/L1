package main

import (
	"strings"
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
	// I guess this is supposed to mean "make a string of length 1<<10"
	// then this might take a long time to build
	// also it would take up more memory than needed - we only
	// use the first 100 symbols!
	v := createHugeString(1 << 10)
	// this is an unnecessary copy - could just assign it right away
	justString = v[:100]
}

// perhaps, something like this would be a better version
func correctedFunc() {
	justString = createHugeString(100)
}

func main() {
	someFunc()
}
