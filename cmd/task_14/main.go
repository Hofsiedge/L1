package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной:
	int, string, bool, channel из переменной типа interface{}
*/

func main() {
	values := []any{10, 0.5, false, make(chan string), struct{}{}}

	fmt.Println("with a type switch:")
	for _, x := range values {
		withTypeSwitch(x)
	}

	fmt.Println("\nwith reflect:")
	for _, x := range values {
		withReflect(x)
	}
}

// withTypeSwitch just uses a type switch, but can't handle unexpected types.
// also, channel types have to be specific, which isn't convenient
func withTypeSwitch(x any) {
	switch typed := x.(type) {
	case int:
		fmt.Printf("%d is an int\n", typed)
	case string:
		fmt.Printf("%q is a string\n", typed)
	case bool:
		fmt.Printf("%v is a bool\n", typed)
	case chan int:
		fmt.Printf("%v is an int channel\n", typed)
	case chan string:
		fmt.Printf("%v is a string channel\n", typed)
	case chan bool:
		fmt.Printf("%v is a bool channel\n", typed)
	case chan struct{}:
		fmt.Printf("%v is an empty struct channel\n", typed)
	default:
		fmt.Printf("could not determine the type of %v (%T). try using reflect\n", x, x)
	}
}

// withReflect determines the type using `reflect` package, so it is not
// performant
func withReflect(x any) {
	kind := reflect.TypeOf(x).Kind()
	switch kind {
	case reflect.Int:
		fmt.Printf("%v is an int\n", x)
	case reflect.String:
		fmt.Printf("%v is a string\n", x)
	case reflect.Bool:
		fmt.Printf("%v is a bool\n", x)
	case reflect.Chan:
		fmt.Printf("%v is a channel\n", x)
	default:
		fmt.Printf("got an unexpected type, but determined it: %v\n", kind.String())
	}
}
