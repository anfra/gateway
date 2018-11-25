package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

var interval int = 1000

type sample struct {
	Time  int64   `json:"time"`
	Id    string  `json:"id"`
	Value float64 `json:"value"`
}

func getSample() sample {
	currentValue := rand.Float64()
	id := "voltage"
	currentSample := sample{time.Now().UTC().Unix(), id, currentValue}
	return currentSample
}

func main() {
	// connect to this socket
	conn, err := net.Dial("tcp", os.Getenv("SERVER_ADDRESS"))
	fmt.Println(conn.RemoteAddr().String())
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// create random values for single metric
		sample := getSample()

		// Encode as JSON
		encoder := json.NewEncoder(conn)
		err := encoder.Encode(sample)
		if err != nil {
			fmt.Println("encode.Encode error: ", err)
		}

		// sleep some time
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
