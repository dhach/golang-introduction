package main

import (
	"fmt"
	"log"
	"strconv"
)

/*
	There is no error handling like try/except or try/catch in Go.
	Instead, it is common to have functions return a value and an error.

	To catch errors, check if the returned error is nil.
	If it is not nil, there has been an error and you can doSomething()
*/

func convertStringToInt(stringToConvert string) int {
	var converted int
	var conversionError error

	// strconv.Atoi returns the desired integer and an error, if it could not convert the input value
	converted, conversionError = strconv.Atoi(stringToConvert)

	// check if the second return value of Atoi is not nil.
	// if true, there has been an error returned
	if conversionError != nil {
		fmt.Println("Error converting string to int!")
		log.Fatal(conversionError) // exits the programm
	}

	return converted
}

/*
	To reduce redundancy and re-use of code snippets,
	it is advisable to create helper functions which help handle errors.

	See: 99_helpers.ho

*/
