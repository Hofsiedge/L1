package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout
*/

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	fmt.Println("with just a wait group")
	justWaitGroup(numbers)

	fmt.Println("with a worker pool")
	withWorkerPool(numbers, 2)
}

// printer starts a goroutine that prints values from an input channel
// and returns a `done` channel
func printer[T any](values <-chan T) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		for s := range values {
			fmt.Println(s)
		}
		// sending 0 bytes
		done <- struct{}{}
	}()
	return done
}

// justWaitGroup combines synchronization a wait group (for workers)
// and a plain channel (for printing)
func justWaitGroup(numbers []int) {
	// I use wg to wait for all the goroutines to finish
	var wg sync.WaitGroup

	// a channel for printing
	results := make(chan int)
	donePrinting := printer(results)
	defer func() {
		// print everything before quiting
		close(results)
		<-donePrinting
	}()

	// set a number of goroutines to wait for
	wg.Add(len(numbers))
	for _, x := range numbers {
		// the variable x is captured as a function argument
		go func(number int) {
			results <- number * number
			wg.Done()
		}(x)
	}
	// don't quit before other goroutines
	wg.Wait()
}

// withWorkerPool shows an example of using a combination of
// worker pool, wait group, done channel
func withWorkerPool(numbers []int, poolSize uint) {
	nums, squares := make(chan int), make(chan int)

	done := printer(squares)
	defer func() {
		// wait for everything to print
		close(squares)
		<-done
	}()
	var wg sync.WaitGroup

	for i := uint(0); i < poolSize; i++ {
		wg.Add(1)
		go func() {
			for x := range nums {
				squares <- x * x
			}
			wg.Done()
		}()
	}

	for _, x := range numbers {
		nums <- x
	}
	close(nums)
	wg.Wait()
}
