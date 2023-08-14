package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Hofsiedge/L1/internal/quicksort"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

// main reads an int slice from command line arguments and writes sorted result
func main() {
	numbers := make([]int, len(os.Args)-1)
	for i, arg := range os.Args[1:] {
		x, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("could not convert argument #%d (%q) to int\n", i, arg)
		}
		numbers[i] = x
	}

	// comparator for stable ascending int sort
	var cmp quicksort.Comparator[int] = func(a, b int) int {
		if a < b {
			return 1
		}
		if a == b {
			return 0
		}
		return -1
	}
	quicksort.Quicksort(numbers, cmp)
	fmt.Println("sorted:")
	fmt.Println(numbers)
}
