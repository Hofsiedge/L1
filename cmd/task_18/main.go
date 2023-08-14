package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться
в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

// Counter embeds sync.Mutex to implement Lock/Unlock
type Counter struct {
	sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.Lock()
	c.value++
	c.Unlock()
}

func (c *Counter) GetValue() int {
	return c.value
}

const (
	poolSize = 4
	seconds  = 5
)

func main() {
	counter := new(Counter)
	var wg sync.WaitGroup
	done := make(chan struct{})
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sum := 0
			for {
				select {
				case <-done:
					fmt.Printf("goroutine %d incremented counter %d times\n",
						id, sum)
					return
				default:
					counter.Increment()
					sum++
				}
			}
		}(i)
	}
	time.Sleep(seconds * time.Second)
	close(done)
	wg.Wait()
	fmt.Printf("final counter value: %d\n", counter.GetValue())
}
