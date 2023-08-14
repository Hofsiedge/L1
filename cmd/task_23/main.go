package main

import (
	"errors"
	"fmt"
)

/*
Удалить i-ый элемент из слайса.
*/

var ErrIndexOutOfBound = errors.New("index out of bounds")

// should only be used to modify a slice, not to make a copy.
// The resulting slice uses the same underlying array as the input slice.
//
// x, ... := RemoveAtIndexUnsafe(x, ...)
func RemoveAtIndexUnsafe[T any](slice []T, index int) ([]T, error) {
	if index < 0 || index >= len(slice) {
		return nil, ErrIndexOutOfBound
	}
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1], nil
}

// RemoveAtIndex is safe to use since it allocates a new array
//
// Note that it still is a shallow copy!
func RemoveAtIndex[T any](slice []T, index int) ([]T, error) {
	if index < 0 || index >= len(slice) {
		return nil, ErrIndexOutOfBound
	}
	result := make([]T, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)
	return result, nil
}

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5}

	// modifying `safe` won't affect `numbers`
	safe, _ := RemoveAtIndex(numbers, 2)
	fmt.Println("safe:", safe, numbers)

	removed, _ := RemoveAtIndexUnsafe(numbers, 2)
	// note that `numbers` is modified as well
	fmt.Println("unsafe:", removed, numbers) // [0 1 3 4 5] [0 1 3 4 5 5]
}
