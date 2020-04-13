package main

import (
	"log"
	// use a local module-level import:
	tcpClient "mytcp/client"
	"net"
)

var (
	targetAddress string
	targetServer  string = "docker-machine"
	targetPort    string = "3333"
	concurrency   int    = 10
)

func main() {
	targetAddress = net.JoinHostPort(targetServer, targetPort)

	// create a new TCP Client and connect to the server
	c := tcpClient.NewClient()
	c.ConnectToServer(targetAddress)

	// create a channel so our goroutines can signal they are done
	doneSignal := make(chan string, concurrency)

	for i := 1; i < concurrency; i++ {
		go func() {
			c.Send("Hello! ") // all goroutines share the same client
			doneSignal <- "payload sent"
		}()
	}

	/*
		This keeps the program running until all goroutines are finished.
		If we would not do this, the program would exit,
		before any goroutine could send data
	*/
	for concurrency > 1 {
		d := <-doneSignal
		log.Printf(d)
		concurrency--
	}
	// the server expects a newline character, so send it after everything has finished
	c.Send("\n")

	c.Close()
}
