package pipeline

import "errors"

// I - input type, O - output type, A & B - intermediate types
type (
	// Producer generates values from an input and feeds them to a Processor
	//
	// Predefined producers: UnfoldProducer
	Producer[I, O any] func(I, chan<- O)

	// Processor processes values from an input channel and feeds them to other
	// processors or to a Consumer.
	//
	// Processor does not have to be a 1 to 1 operation - e.g., it might filter
	// out unnecessary values or generate multiple values from one.
	//
	// Predefined processors: MapProcessor, EffectProcessor, FilterProcessor
	Processor[I, O any] func(<-chan I) <-chan O

	// Consumer receives values from a channel and outputs a single result after
	// the input channel is closed
	//
	// Predefined consumers: StubConsumer, EffectConsumer, ReduceConsumer
	Consumer[I, O any] func(<-chan I) O

	Pipeline[I, A, B, O any] struct {
		prod Producer[I, A]
		proc Processor[A, B]
		cons Consumer[B, O]

		in      chan A
		running bool
		result  chan O
	}
)

// MapProcessor applies a function to every value in the
// input channel and writes every result to the output channel
func MapProcessor[I, O any](f func(I) O) Processor[I, O] {
	return func(in <-chan I) <-chan O {
		out := make(chan O)
		go func() {
			for v := range in {
				out <- f(v)
			}
		}()
		return out
	}
}

// EffectProcessor applies an effect to every value from the
// input channel and writes the value to the output channel unchanged
func EffectProcessor[T any](effect func(T)) Processor[T, T] {
	return func(in <-chan T) <-chan T {
		out := make(chan T)
		go func() {
			for v := range in {
				effect(v)
				out <- v
			}
		}()
		return out
	}
}

// FilterProcessor applies a predicate to input values and only writes
// those that returned true
func FilterProcessor[T any](predicate func(T) bool) Processor[T, T] {
	return func(in <-chan T) <-chan T {
		out := make(chan T)
		go func() {
			for v := range in {
				if predicate(v) {
					out <- v
				}
			}
		}()
		return out
	}
}

// UnfoldProducer writes every value from the source slice to the output channel
func UnfoldProducer[T any]() Producer[[]T, T] {
	return func(source []T, out chan<- T) {
		for _, v := range source {
			out <- v
		}
	}
}

// StubConsumer is the most trivial consumer - it does nothing to the inputs
// and returns an empty struct value
func StubConsumer[T any]() Consumer[T, struct{}] {
	return func(in <-chan T) struct{} {
		for range in {
		}
		return struct{}{}
	}
}

// EffectConsumer applies an effect to inputs and return an empty struct value
func EffectConsumer[T any](effect func(T)) Consumer[T, struct{}] {
	return func(in <-chan T) struct{} {
		for x := range in {
			effect(x)
		}
		return struct{}{}
	}
}

// ReduceConsumer reduces elements of input channel with the provided function,
// returns result.
//
// neutralElement is the first value of the first argument to reduce.
// then the result of previous invocation is used as first argument.
//
// examples of usage: sum of inputs; unique string set constructor; etc
func ReduceConsumer[I, O any](reduce func(O, I) O, neutralElement O) Consumer[I, O] {
	return func(in <-chan I) O {
		value := neutralElement
		for x := range in {
			value = reduce(value, x)
		}
		return value
	}
}

var (
	ErrPipelineIsAlreadyRunning = errors.New("pipeline is already running")
	ErrPipelineIsNotRunning     = errors.New("pipeline is not running")
)

// NewPipeline initializes a new Pipeline
//
// Only one processor is supported. To use multiple processors chain them with Chain(proc1, proc2)
func NewPipeline[I, A, B, O any](
	prod Producer[I, A], proc Processor[A, B], cons Consumer[B, O],
) *Pipeline[I, A, B, O] {
	pipeline := Pipeline[I, A, B, O]{
		prod:    prod,
		proc:    proc,
		cons:    cons,
		in:      nil,
		running: false,
		result:  nil,
	}
	return &pipeline
}

// Run initializes channels and starts the processor and the consumer.
//
// Run can be run multiple times on the same Pipeline but not when the
// Pipeline is already running.
//
// Run is not blocking
func (p *Pipeline[I, A, B, O]) Run() error {
	if p.running {
		return ErrPipelineIsAlreadyRunning
	}
	p.in = make(chan A)
	p.result = make(chan O, 1)
	p.running = true

	m := p.proc(p.in)
	go func() {
		p.result <- p.cons(m)
	}()
	return nil
}

// Feed calls the producer to feed new values into the pipeline
//
// Feed is blocking
func (p *Pipeline[I, A, B, O]) Feed(source I) error {
	if !p.running {
		return ErrPipelineIsNotRunning
	}
	p.prod(source, p.in)
	return nil
}

// Close closes all the channels and return the result of computations
//
// Close is blocking
func (p *Pipeline[I, A, B, O]) Close() (O, error) {
	if !p.running {
		var zero O
		return zero, ErrPipelineIsNotRunning
	}
	defer close(p.result)
	close(p.in)
	return <-p.result, nil
}

// Chain joins two sequential processors into an equivalent one
func Chain[I, M, O any](a Processor[I, M], b Processor[M, O]) Processor[I, O] {
	return func(in <-chan I) <-chan O {
		m := a(in)
		return b(m)
	}
}
