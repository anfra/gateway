package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

var interval int = 1

type sample struct {
	time  time.Time
	id    string
	value int
}

func main() {
	// connect to this socket
	conn, _ := net.Dial("tcp", os.Getenv("SERVER_ADDRESS"))
	fmt.Println(conn.RemoteAddr().String())
	for {

		// create random values for single metric

		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		// send to socket
		fmt.Fprintf(conn, text+"\n")

		time.Sleep(time.Duration(interval) * time.Second)
	}
}
