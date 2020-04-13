package main

import (
	"fmt"
)

func variablesBasics() {
	printTraceAndGap()

	/*
	 Declaring and initializing
	*/
	// most complete version:
	var myStatus string
	myStatus = "presenting"

	// more concise
	var myNewStatus string = "still presenting"

	// without specified type
	myNewerStatus := "still presenting, duh"

	/*
		It is an error to...
		- declare and not use a variable
		- redeclare a variable
		- mismatch the type
	*/

	fmt.Println(myStatus)
	fmt.Println(myNewStatus)
	fmt.Println(myNewerStatus)

}

/*
	basic data types:
		- bool
		- string
		- int8, int16, int32, int64
		- uint8, uint16, uint32, uint64
		- float32, float64
		- byte

	special data types:
		- complex64, complex128
		- rune


	A byte is actually an alias for uint8
	A rune is actually an alias for int32

	int, unint are 32 bits wide on 32-bit systems, 64 bits wide on 64-bit systems
*/
