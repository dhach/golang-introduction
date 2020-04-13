package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// Use io/ioutil helper function
func readFileViaIoutil() {
	printTraceAndGap()

	data, err := ioutil.ReadFile(fileToRead)

	if checkError(err) {
		fmt.Println("Error reading file!")
		os.Exit(1)
	}

	fmt.Print(string(data)) // data is of type 'byte'
}

// open a file directly and use bufio scanner
func readFileViaBuffer() {
	printTraceAndGap()

	fileHandle, err := os.Open(fileToRead)
	if checkError(err) {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}
	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

/*







 */
// Use io/ioutil helper function
func readFile() []byte {

	data, err := ioutil.ReadFile(fileToRead)

	if checkError(err) {
		fmt.Println("Error reading file!")
		os.Exit(1)
	}

	return data
}
