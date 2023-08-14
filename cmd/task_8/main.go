package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Дана переменная int64.
Разработать программу, которая устанавливает i-й бит в 1 или 0.
*/

/*
this package has a test
*/

// sentient errors
var (
	ErrInvalidBit          = errors.New("bit position is out of bounds")
	ErrWrongArgumentNumber = errors.New("unexpected number of arguments")
)

func main() {
	// check the number of arguments
	if len(os.Args) != 4 {
		err := fmt.Errorf("%w: %d", ErrWrongArgumentNumber, len(os.Args)-1)
		fmt.Println(err)
		return
	}
	// parse the arguments
	value, valueErr := strconv.ParseInt(os.Args[1], 10, 64)
	bit, bitErr := strconv.Atoi(os.Args[2])
	one, oneErr := strconv.ParseBool(os.Args[3])
	// if any of the errors are not nil
	if valueErr != nil || bitErr != nil || oneErr != nil {
		invalidArgs := make([]string, 0)
		if valueErr != nil {
			invalidArgs = append(invalidArgs,
				fmt.Sprintf("1 - expected int64, got %q;", os.Args[1]))
		}
		if bitErr != nil {
			invalidArgs = append(invalidArgs,
				fmt.Sprintf("2 - expected int, got %q;", os.Args[2]))
		}
		if oneErr != nil {
			invalidArgs = append(invalidArgs,
				fmt.Sprintf("3 - expected bool, got %q;", os.Args[3]))
		}
		lines := strings.Join(invalidArgs, "\n")
		// showing an error message
		fmt.Printf("unexpected format for the following arguments: %s\n", lines)
		return
	}
	result, err := setBit(value, bit, one)
	// if there was an error, print it and exit
	if err != nil {
		fmt.Println(err)
		return
	}
	// print the results if everything's OK
	fmt.Println(result)
}

func setBit(value int64, bit int, one bool) (int64, error) {
	// check the bit range (int64 has 64 bits)
	if bit < 0 || bit >= 64 {
		// returning a sentient error
		return 0, fmt.Errorf("%w: %d", ErrInvalidBit, bit)
	}
	// bit shifting a one to place it at the needed bit
	var i int64 = 1 << bit
	// if want to set a one - or
	if one {
		return value | i, nil
	}
	// if want to set a zero - bit clear
	return value &^ i, nil
}
