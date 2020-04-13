# Examples

## TCP server & client

This example consists of a simple TCP server and TCP client.

The server (directoy "tcp-server") listens on port 3333 and is asynchronous. It can handle multiple connections and prints the received string to stdout. It will buffer the received data until it receives a newline charachter ('\n')
The Makefile will build the server inside a Docker container and also start it in one.

The client is built and run on the local machine.
It will connect to the server and send a predefined string. The string will be sent for a predefined amount of times using goroutines and basic channel signalling.

### Go Modules

Note that the TCP client is containing a file called "go.mod".
This essentially tells Golang that all Go-files inside this directory, as well as everything below it, belong to a specific, named package and thus allows you to do local-level imports. It can also be used to write applications for importing.
