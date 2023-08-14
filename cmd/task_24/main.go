package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными
параметрами x, y и конструктором.
*/

/*
I implemented 2 distance functions - l1 & l2
*/

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type Point[T Number] struct {
	x, y T
}

// Stringer
func (p Point[T]) String() string {
	return fmt.Sprintf("{%v, %v}", p.x, p.y)
}

// constructor
func NewPoint[T Number](x, y T) *Point[T] {
	return &Point[T]{x, y}
}

// L1 (Manhattan) distance
func L1Distance[T Number](a, b *Point[T]) float64 {
	deltaX, deltaY := float64(a.x-b.x), float64(a.y-b.y)
	return math.Abs(deltaX) + math.Abs(deltaY)
}

// L2 (Euclidean) distance
func L2Distance[T Number](a, b *Point[T]) float64 {
	deltaX, deltaY := float64(a.x-b.x), float64(a.y-b.y)
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func main() {
	A := NewPoint(0.0, 5.0)
	B := NewPoint(3.0, 1.0)
	fmt.Printf("l1(%v, %v) = %v\n", A, B, L1Distance(A, B))
	fmt.Printf("l2(%v, %v) = %v\n", A, B, L2Distance(A, B))
}
