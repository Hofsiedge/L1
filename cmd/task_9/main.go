package main

import (
	"fmt"
	"math/rand"

	"github.com/Hofsiedge/L1/internal/pipeline"
)

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/

func main() {
	var numbers [100]int
	for i := 0; i < 100; i++ {
		numbers[i] = rand.Int() % 100
	}

	fmt.Println("with a generic pipeline:")
	withGenericPipeline(numbers[:])

	fmt.Println("without the overkill abstraction:")
	withSimplePipeline(numbers[:])
}

func withGenericPipeline(numbers []int) {
	// prod writes the numbers into a channel
	prod := pipeline.UnfoldProducer[int]()
	// proc squares the numbers and then prints them
	square := pipeline.MapProcessor(func(x int) int { return x * x })
	// cons just prints the inputs
	cons := pipeline.EffectConsumer[int](func(x int) {
		fmt.Println(x)
	})

	// this pipeline takes an int slice, feeds ints into a proccessor
	// that turns them into ints that are consumed by the consumer with
	// empty return value
	pipe := pipeline.NewPipeline(prod, square, cons)
	pipe.Run()
	// Feed is blocking and can be run multiple times
	// could be run from a goroutine to parallelize consumption
	pipe.Feed(numbers[:len(numbers)/2])
	pipe.Feed(numbers[len(numbers)/2:])
	// Close returns the result of computation and closes the pipeline
	// the error is ignored here because I know for sure that it will be nil
	// the values is ignored since it is just struct{}{}
	_, _ = pipe.Close()
}

// same as withGenericPipeline, but without abstractions
func withSimplePipeline(numbers []int) {
	in := make(chan int)
	out := squarer(in)

	done := make(chan struct{})
	// print results
	go func() {
		for x := range out {
			fmt.Println(x)
		}
		done <- struct{}{}
	}()

	// inputs
	for _, x := range numbers {
		in <- x
	}
	close(in)
	<-done
}

func squarer(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for x := range in {
			out <- x * x
		}
		close(out)
	}()
	return out
}
