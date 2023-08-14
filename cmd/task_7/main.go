package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

func main() {
	names := []string{"A", "B", "C"}
	mapSize := 1000
	withMutex(names, mapSize)
	withSyncMap(names, mapSize)
}

type LockableMap[K comparable, V any] struct {
	sync.Mutex
	Map map[K]V
}

// withMutex uses a mutex to prevent concurrent map writing
// withMutex uses a `LockableMap` struct that embeds `sync.Mutex`
// but the mutex could be a separate value
func withMutex(names []string, mapSize int) {
	fmt.Println("with a mutex and a map:")
	m := LockableMap[int, string]{
		Map: make(map[int]string),
	}

	var wg sync.WaitGroup
	keys := make(chan int)

	for _, value := range names {
		wg.Add(1)
		go func(v string) {
			count := 0
			for k := range keys {
				m.Lock()
				m.Map[k] = v
				m.Unlock()
				count++
			}
			fmt.Printf("%s set %d values\n", v, count)
			wg.Done()
		}(value)
	}

	for i := 0; i < mapSize; i++ {
		keys <- i
	}
	close(keys)
	wg.Wait()
}

// withSyncMap does not need a separate mutex - it is embedded in
// `sync.Map` and is handled by it
func withSyncMap(names []string, mapSize int) {
	fmt.Println("with a sync.Map:")
	m := sync.Map{}
	var wg sync.WaitGroup
	keys := make(chan int)

	for _, value := range names {
		wg.Add(1)
		go func(v string) {
			count := 0
			for k := range keys {
				m.Store(k, v)
				count++
			}
			fmt.Printf("%s set %d values\n", v, count)
			wg.Done()
		}(value)
	}

	for i := 0; i < mapSize; i++ {
		keys <- i
	}
	close(keys)
	wg.Wait()
}
