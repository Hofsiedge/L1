package main

import (
	"fmt"

	"github.com/Hofsiedge/L1/internal/bisect"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	numbers := []int{-20, -3, 10, 100}
	var comparator bisect.Comparator[int] = func(a, b int) bisect.ComparatorOrder {
		if a < b {
			return bisect.OrderedRight
		}
		if a == b {
			return bisect.OrderedEqual
		}
		return bisect.OrderedWrong
	}
	fmt.Println(numbers)
	for _, value := range []int{0, -3, 150} {
		position, found := bisect.Bisect(numbers, comparator, value)
		fmt.Printf("%v: %v @ %d\n", value, found, position)
	}
}
