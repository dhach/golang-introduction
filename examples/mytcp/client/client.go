package client

import (
	"bufio"
	"log"
	"net"
	"sync"
)

// TCPClient is a basic client for sending data over TCP
type TCPClient struct {
	conn   net.Conn
	writer *bufio.Writer
	sync.Mutex
	/*
		Mutex ensures there is a lock on each TCPClient,
		so that bufio.Writer remains in a clean state.
		This is done by "Locking" and "Unlocking" it in the Send() method
	*/
}

// NewClient returns an instance of TCPClient
// we choose this approach so we can have an "uninitialized" instance and use methods to configure it
func NewClient() *TCPClient {
	return &TCPClient{}
}

// ConnectToServer establishes a connection to the target server.
func (c *TCPClient) ConnectToServer(targetAddress string) {
	conn, err := net.Dial("tcp", targetAddress)
	if err != nil {
		log.Fatal("Cannot connect to server ", targetAddress)
	}
	c.conn = conn
	c.writer = bufio.NewWriter(conn)
}

// Send sends a TCP data payload to the target server
func (c *TCPClient) Send(payload string) {
	c.Lock() // acquire the Mutex lock to keep the Writer in a clean state
	_, err := c.writer.WriteString(payload)
	if err == nil {
		err = c.writer.Flush()
		if err != nil {
			log.Fatal("Cannot flush payload to server")
		}
	} else {
		log.Fatal("Cannot send payload to server")
	}

	c.Unlock() // release the lock
}

// Close disconnects from the server
func (c *TCPClient) Close() {
	c.Lock()
	defer c.Unlock()
	c.conn.Close()
}
