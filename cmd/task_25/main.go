package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Hofsiedge/L1/internal/sleep"
)

/*
Реализовать собственную функцию sleep.
*/

/*
the original implementation (src/runtime/time.go in the Go source code)
uses gopark which is private.
So these are the only ways I could come up with on spot.
*/

// AfterSleep blocks until the `duration` time is out.
// feels like cheating - time.After is basically the same as time.Sleep
func AfterSleep(duration time.Duration) {
	<-time.After(duration)
}

// BraindeadSleep is constantly checking for the time,
// returns when current time is past the end time.
// could be modified to use time.Ticker to reduce CPU load
func BraindeadSleep(duration time.Duration) {
	endTime := time.Now().Add(duration)
	for time.Now().Before(endTime) {
	}
}

func main() {
	var wg sync.WaitGroup

	// this sleeper uses `alarm` syscall
	sleeper := sleep.NewSleeper()
	defer sleeper.Close()

	for _, seconds := range []int{10, 4, 6, 15} {
		wg.Add(1)
		go func(s int) {
			defer wg.Done()
			sleeper.Sleep(s)
			fmt.Println(s)
		}(seconds)
	}
	wg.Wait()
}
