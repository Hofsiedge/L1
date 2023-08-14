package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree).
Создать для нее собственное множество.
*/

// this is implemented in the most straight-forward way
//
// could be implemented with Set[T] from task 11 as well
func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})
	for _, str := range input {
		set[str] = struct{}{}
	}

	fmt.Println(set)
}
