package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// Use io/ioutil helper function
func writeFileViaIoutil() {
	printTraceAndGap()

	var dataToWrite []byte
	dataToWrite = readFile()

	err := ioutil.WriteFile(fileToWrite, dataToWrite, 0644)
	if checkError(err) {
		fmt.Println("Error writing to file: ", err)
		os.Exit(2)
	}

}

// open a file directly and use bufio scanner
func writeFileViaBufio() {
	printTraceAndGap()

	var dataToWrite []byte
	dataToWrite = readFile()

	fileHandle, errOpen := os.Open(fileToWrite)
	if checkError(errOpen) {
		fmt.Println("Error writing to file: ", errOpen)
		os.Exit(2)
	}

	writer := bufio.NewWriter(fileHandle)
	_, errWrite := writer.Write(dataToWrite)

	if checkError(errWrite) {
		fmt.Println("Error writing to file: ", errOpen)
		os.Exit(2)
	}
}
