package main

import (
	"fmt"
)

func arrays() {
	printTraceAndGap()
	/*
		Arrays:
			- fixed-length
			- typed
	*/

	var stringArr [4]string // an array of strings, length 4
	stringArr[0] = "Hello"
	stringArr[1] = ","
	stringArr[2] = "dear"
	stringArr[3] = "Sir/Madam"

	// it is an error to write beyound the given length:
	// stringArr[4] = "Dummy"

	// iterate over the array:
	for i, v := range stringArr {
		fmt.Printf("[%d] %s\n", i, v)
	}

	// ntk: indexing a string yields its bytes, not its char:
	fmt.Println("byte at position [0][0] ", stringArr[0][0])

}

func slices() {
	printTraceAndGap()
	/*
		Slices:
			- like arrays, but without fixed length
	*/

	var intSlice []int

	// it is an error to write to a slice if the index is not yet "occupied"
	// intSlice[0] = 100
	// instead, append to the slice:
	intSlice = append(intSlice, 100)
	intSlice = append(intSlice, 200)

	// then we can assign a value directly to an index
	intSlice[0] = 300

	for _, v := range intSlice {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

}

func maps() {
	printTraceAndGap()
	/*
		Maps:
			- arrays with named indices
			- aka 'Dictionaries'

		!!! Caution !!!
		They have to be initialized.
		Simply declaring a map will cause it to be a nil pointer, to which cannot be written!
	*/

	var numbersToText map[int]string // KeyType: int, ValueType: string

	// numbersToText[10] = "Ten" // <-- this will crash with an error: "panic: assignment to entry in nil map"

	numbersToText = make(map[int]string) // initialize the map

	numbersToText[1] = "One"
	numbersToText[10] = "Ten"
	numbersToText[25] = "Twentyfive"

	delete(numbersToText, 25) // deletes in-place

	fmt.Println(numbersToText[10])

	// test for existence of a key and get its value:
	v, succ := numbersToText[1]
	if succ {
		fmt.Println(v)
	} else {
		fmt.Println("KeyError")
	}

	// iterate over all entries in the map:
	for key, value := range numbersToText {
		fmt.Printf("%d: %s\n", key, value)
	}

}
