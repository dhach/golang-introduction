package main

import (
	"fmt"
	"math/rand"
)

func loopExamples() {
	intArray := [5]int{10, 20, 30, 40, 50}

	/*
		classic-style loop over an array
	*/
	for i := 0; i < len(intArray); i++ {
		fmt.Printf("%d ", intArray[i])
	}
	fmt.Printf("\n\n")

	strSlice := []string{"Hello", ",", "World", ".", "You", "kinda", "suck", "sometimes"}

	/*
		using 'range' function
	*/
	for index, elem := range strSlice {
		fmt.Printf("index %d: %v\n", index, elem)
	}

	/*

		There are no 'while' loops in Go.
		Only 'for':
	*/
	i := 1
	for i <= 10 {
		fmt.Printf("%d ", i)
		fmt.Printf("\n")
		i++
	}

	// 'while true' with exit condition:
	cond := false
	for {
		fmt.Println("while loop")
		if !cond {
			break
		}
	}

	fmt.Print("\n\n\n")
}

/*
	conditionals are similiar to C/C++
*/
func conditions() {
	// simple if
	if true {
		fmt.Println("true")
	}

	// comple if/elif/else
	if !true {
		fmt.Println("false")
	} else if false {
		fmt.Println("true")
	} else {
		fmt.Println("else")
	}

	fmt.Print("\n\n\n")
}

/*
	switches are almost like in every other language
*/
func switches() {
	val := 25

	switch val {
	case 10:
		fmt.Println("Ten")
	case 20:
		fmt.Println("Twenty")
	case 30:
		fmt.Println("Thirty")
	default:
		fmt.Println("something else")
	}

	// switches in Go are not limited to integers:
	r := rand.Intn(50)
	switch { // no condition means true
	case r%2 == 0:
		fmt.Println("even")
		fallthrough
	case r%2 == 1:
		fmt.Println("un-even")
		fallthrough
	case r > 25:
		fmt.Println("> 25")
	case r < 25:
		fmt.Println("< 25")
	default:
		fmt.Println("This part should never be reached!")
		panic("THE END IS NIGH!")
	}

	fmt.Print("\n\n\n")
}
