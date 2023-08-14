package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	withRange()
	withBreakOnNotOk()
	withSeparateChannel()
	withCancel()

	// 5. just quit main
	// defers are not executed which might be a problem
	// if there are resources needing cleaning up
	go func() {
		defer fmt.Println("defer is not executed as well")
		fmt.Println("goroutine 5 will be abruptly closed with main")
		time.Sleep(time.Second * 2)
		fmt.Println("would be printed if main waited for it")
	}()
	time.Sleep(time.Second)
	// exiting main naturally or with `os.Exit`
	// results in abruptly closing running goroutines
}

// 1. Use `for ... := range ch` and `close(ch)`
func withRange() {
	// for synchronization
	var wg sync.WaitGroup

	ch := make(chan string)
	wg.Add(1)
	go func() {
		for value := range ch {
			fmt.Println(value)
		}
		fmt.Println("goroutine 1 is stopped!")
		wg.Done()
	}()
	ch <- "goroutine 1: A"
	ch <- "goroutine 1: B"
	// closing the channel will break the for loop in
	// all listening goroutines
	// so it can be used to stop multiple goroutines
	close(ch)
	wg.Wait()
}

// 2. Use `for {}`, `value, ok := <-ch`, `break`
func withBreakOnNotOk() {
	// for synchronization
	var wg sync.WaitGroup

	ch := make(chan string)
	wg.Add(1)
	go func() {
		for {
			value, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(value)
		}
		fmt.Println("goroutine 2 is stopped!")
		wg.Done()
	}()
	ch <- "goroutine 2: A"
	ch <- "goroutine 2: B"
	// all the listening goroutines will get
	// a (zero value, false) pair
	// so it can be used to stop multiple goroutines
	close(ch)
	wg.Wait()
}

// 3. Use a separate channel and `select` statement
func withSeparateChannel() {
	// for synchronization
	var wg sync.WaitGroup

	ch := make(chan string)
	done := make(chan struct{})
	wg.Add(1)
	go func() {
	loop:
		for {
			select {
			case <-done:
				break loop
			case value, ok := <-ch:
				if !ok {
					break loop
				}
				fmt.Println(value)
			}
		}
		fmt.Println("goroutine 3 is stopped!")
		wg.Done()
	}()
	ch <- "goroutine 3: A"
	ch <- "goroutine 3: B"
	// closing a channel makes *all* the goroutines listening with
	// `v, ok <- done` receive operation receive  a (zero value, false) pair.
	// so it can be used to close multiple goroutines
	close(done)
	wg.Wait()
	close(ch)
}

// 4. Same as 3, but using context.WithCancel instead of a raw channel
func withCancel() {
	// for synchronization
	var wg sync.WaitGroup

	ch := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			case value, ok := <-ch:
				if !ok {
					break loop
				}
				fmt.Println(value)
			}
		}
		fmt.Println("goroutine 4 is stopped!")
		wg.Done()
	}()
	ch <- "goroutine 4: A"
	ch <- "goroutine 4: B"

	// calling `cancel` closes context's underlying `done` channel
	// (called `c` in the `context` source code)
	cancel()
	wg.Wait()
	close(ch)
}
