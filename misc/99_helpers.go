package main

import (
	"log"
	"runtime"
)

// checkError checks if the passed error is nil. It returns true if there is an error and false if not
func checkError(err error) bool {
	if err != nil {
		return true
	}
	return false
}

// printTraceAndGap is a helper function to simply print the current function and some newlines
func printTraceAndGap() {
	log.Print("\n\n\n")
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		log.Println("?", 0, "?")
	}
	fn := runtime.FuncForPC(pc)
	log.Printf("%s: %d  [ %s ]\n\n", file, line, fn.Name())
}
