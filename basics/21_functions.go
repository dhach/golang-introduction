package main

import (
	"fmt"
)

/*
	functions are much the same as in any other programming language

*/

// simplest form: no argument, no return value
func basicFuntionOne() {
	fmt.Println("Basic function #1")
}

// with argument
func basicFunctionTwo(pronoun string) {
	fmt.Println("Basic function #2")
	fmt.Println("Hello, ", pronoun)
}

// with argument and return value
func functionThree(pronoun string) bool {
	var success bool = true

	fmt.Println("Basic function #3")
	fmt.Println("Hello, ", pronoun)

	return success
}

// with argument and named return value
func functionFour(pronoun string) (success bool) {
	success = false // always assign a default value to named return values

	fmt.Println("function #4")
	fmt.Println("Hello, ", pronoun)

	success = true

	return
}
