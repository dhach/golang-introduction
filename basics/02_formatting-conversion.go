package main

import (
	"fmt"
	"strconv"
)

/*
	Variables can be declared outside of functions.
	These are globally scoped.

	You can also use it inside a "var block", like with imports.
*/
var (
	name         string
	mood         string
	age          uint8
	strength     float64
	consentGiven bool
)

func stringFormatting() {
	printTraceAndGap()
	/*
		Basic string formatting.
		Behaves like any other string formatting, like in C or Python
	*/

	name = "James Howlett"
	mood = "angry"
	age = 130
	consentGiven = false
	strength = 999.99

	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Mood: %s\n", mood)
	fmt.Printf("Strength: %f\n", strength)
	fmt.Printf("Consent: %t\n", consentGiven)
}

func typeConversion() {
	printTraceAndGap()
	/*
		Basic type conversion.
		Either directly via casting or via functions
	*/

	var ageFloat float64 = float64(age)
	fmt.Println(ageFloat)

	var ageIntAgain int = int(ageFloat)
	fmt.Println(ageIntAgain)

	/*
		Casting to a string or back is more complex.
		It requires usage of a special function from the package "strconv"
	*/

	// string to int: Atoi
	ageStringFromInt := strconv.Itoa(ageIntAgain)
	ageIntFromString, _ := strconv.Atoi(ageStringFromInt)

	fmt.Println(ageStringFromInt, ageIntFromString)
}
