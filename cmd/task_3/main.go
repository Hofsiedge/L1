package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов с использованием конкурентных вычислений
*/

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	fmt.Printf("with a mutex: %d\n", withMutex(numbers))
	fmt.Printf("with channels: %d\n", withChannels(numbers, 2))
}

// withMutex uses a mutex to synchronize sum editing
func withMutex(numbers []int) int {
	// I use wg to wait for all the goroutines to finish
	var wg sync.WaitGroup

	// embedding a mutex to provide the sum with Lock/Unlock methods.
	// mutex could be a separate variable though
	sum := struct {
		sync.Mutex
		sum int
	}{sync.Mutex{}, 0}

	// set a number of goroutines to wait for
	wg.Add(len(numbers))
	for _, x := range numbers {
		// the variable x is captured as a function argument
		go func(number int) {
			square := number * number

			// locking before editing to prevent data races
			sum.Lock()
			sum.sum += square
			// unlocking to make the sum editable again
			sum.Unlock()

			wg.Done()
		}(x)
	}
	// don't quit before other goroutines
	wg.Wait()
	return sum.sum
}

// withChannels synchronizes sum updating with just channels (and a wait group)
func withChannels(numbers []int, poolSize uint) int {
	nums, squares := make(chan int), make(chan int)

	// sum updating goroutine
	done := make(chan int)
	go func() {
		sum := 0
		for s := range squares {
			sum += s
		}
		done <- sum
	}()

	var wg sync.WaitGroup

	// workers
	wg.Add(int(poolSize))
	for i := uint(0); i < poolSize; i++ {
		// could keep a local sum as well to reduce the number of `square <-`
		go func() {
			for x := range nums {
				squares <- x * x
			}
			wg.Done()
		}()
	}

	// sending jobs to workers
	for _, x := range numbers {
		nums <- x
	}
	close(nums)

	// waiting for workers to finish
	wg.Wait()
	// signaling the sum goroutine to stop
	close(squares)
	return <-done
}
