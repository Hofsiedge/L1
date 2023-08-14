package main

/*
Поменять местами два числа без создания временной переменной.
*/

import "fmt"

func main() {
	// the simplest way
	a, b := 10, 15
	a, b = b, a
	fmt.Println(a, b) // 15 10

	// the same, but in a slice
	qux := []int{0, 1, 2, 3, 4}
	qux[0], qux[3] = qux[3], qux[0]
	fmt.Println(qux) // [3 1 2 0 4]

	// using an operation and its inverse
	// with a + and -
	// (only for numbers with a small enough sum to fit in the type)
	c, d := 10, 15
	c += d
	d = c - d
	c = c - d
	fmt.Println(c, d) // 15 10
	// could be any other operation and its inverse (e.g. * and /)
	// but each one has its own restrictions
}
