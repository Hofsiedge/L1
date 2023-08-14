package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает
две числовых переменных a,b, значение которых > 2^20.
*/

func main() {
	A := big.NewInt(1073741824)  // 2^30
	B := big.NewInt(68719476736) // 2^36

	sum := new(big.Int).Add(A, B)
	diff := new(big.Int).Sub(B, A)
	prod := new(big.Int).Mul(A, B)
	quot := new(big.Int).Div(B, A)

	fmt.Printf("%v + %v = %v\n", A, B, sum)
	fmt.Printf("%v - %v = %v\n", B, A, diff)
	fmt.Printf("%v * %v = %v\n", A, B, prod)
	fmt.Printf("%v / %v = %v\n", B, A, quot)
}
