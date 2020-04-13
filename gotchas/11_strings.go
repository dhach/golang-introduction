package main

import (
	"log"
	"strconv"
	"strings"
)

/*

Gotcha #1:	very inefficient string building

*/
func inefficientStringBuilding() {
	printTraceAndGap()

	/*
		Creating huge strings via ' str += "char" ' will result in excessive copying of memory contents into new buffers.
		This can easily result in a run time in the order of minutes, instead of merely a second. True story!
	*/

	var myString string

	for i := 0; i < 10; {
		for i := 1; i < 2000; i++ {
			myString += strconv.Itoa(i)
		}
	}
	log.Println(myString)

}

func efficientStringBuilding() {
	printTraceAndGap()

	/*
		Instead, use StringBuilder from the standard library 'strings'.

		It is pretty powerful, easy to use and really fast in more recent version of Go
	*/

	var myStringBuilder strings.Builder

	for i := 0; i < 10; {
		for i := 1; i < 2000; i++ {
			myStringBuilder.WriteString(strconv.Itoa(i))
		}
	}

	log.Println(myStringBuilder.String())
}
