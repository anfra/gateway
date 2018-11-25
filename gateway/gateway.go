package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

const (
	MyDB     = "HISTORY"
	username = ""
	password = ""
)

type sample struct {
	Time  int64   `json:"time"`
	Id    string  `json:"id"`
	Value float64 `json:"value"`
}

func writePoints(clnt client.Client, sample sample) {
	sampleSize := 1000

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	pt, err := client.NewPoint(
		"cpu_usage",
		tags,
		fields,
		time.Now(),
	)
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	if err := clnt.Write(bp); err != nil {
		log.Fatal(err)
	}
}

func handleServerConnection(conn net.Conn) {

	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     os.Getenv("INFLUX_URL"),
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	for {
		var msg sample
		d := json.NewDecoder(conn)
		err := d.Decode(&msg)
		fmt.Println(msg, err)
	}

	writePoints(c, msg)
}

func main() {

	fmt.Println("Launching gateway...")

	// listen on all interfaces
	ln, err := net.Listen("tcp", ":"+os.Getenv("LISTENER_PORT"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening on ", ln.Addr())

	for {
		// accept connection on port
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(conn)
	}
}
