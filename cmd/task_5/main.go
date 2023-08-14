package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения
в канал, а с другой стороны канала — читать.
По истечении N секунд программа должна завершаться.
*/

var workingTime time.Duration

func main() {
	// setting and reading `workingTime` as a flag
	flag.DurationVar(&workingTime, "time", 5*time.Second, "working time")
	flag.Parse()
	if workingTime < 0 {
		fmt.Printf("`time` can't be negative (got %d)\n", workingTime)
		os.Exit(0)
	}

	// ticker writes to ticker.C every second
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// an `expired` channel & time.AfterFunc could be used
	// instead of time.After(...)
	/*
		// after `workingTime` seconds an expiration message will be sent
		expired := make(chan struct{})
		_ = time.AfterFunc(time.Duration(workingTime)*time.Second, func() {
			expired <- struct{}{}
		})
	*/

	expired := time.After(workingTime)
	messages := make(chan string, 1)
	for {
		select {
		// after `workingTime`
		case <-expired:
			return
		// every second
		case currentTime := <-ticker.C:
			// writing to `messages` will not block since the channel
			// is buffered and has a free spot
			messages <- currentTime.String()
			fmt.Println(<-messages)
		}
	}
}
