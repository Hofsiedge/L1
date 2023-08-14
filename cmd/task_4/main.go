package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из
канала и выводят в stdout. Необходима возможность выбора количества
воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

/*
The workers are reading from the `messages` channel via a range statement.
The main goroutine is writing random ints to `messages` until it
receives an interrupt signal from the channel `sig`.
Then it closes the `messages` channel. When a worker attempts to
read from `messages` its for loop exits.

An alternative solution could be using a select statement in the workers
that would wait for a separate channel closing and write by default.
That separate channel could be encapsulated by context.WithCancel as well.
But both of these options introduce unnecessary entities - the `messages`
channel is the only important channel and signalling with its closing
is enough.
*/

var poolSize uint

func main() {
	// defining and parsing a flag
	flag.UintVar(&poolSize, "n", 2, "worker pool size")
	flag.Parse()

	messages := make(chan int)

	// starting workers
	var wg sync.WaitGroup
	var mx sync.Mutex
	wg.Add(int(poolSize))
	for i := uint(0); i < poolSize; i++ {
		go func(id uint) {
			for message := range messages {
				// use mutex to prevent malformed outout
				mx.Lock()
				fmt.Printf("goroutine %d: %d\n", id, message)
				mx.Unlock()
			}
			wg.Done()
		}(i)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

loop:
	// writing messages
	for {
		select {
		case <-sig:
			// using a `break` statement with a label to break out of `for`,
			// not out of `select`
			break loop
		default:
			messages <- rand.Int()
		}
	}
	// signalling workers to stop
	close(messages)
	// waiting for workers
	wg.Wait()
}
