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
	Time  time.Time `json:"time"`
	Id    string    `json:"id"`
	Value float64   `json:"value"`
}

// queryDB convenience function to query the database
func queryDB(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: MyDB,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

func writePoints(clnt client.Client, msg sample) {

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	tags := map[string]string{
		"voltage": "voltage_1",
	}

	fields := map[string]interface{}{
		"value": msg.Value,
	}

	pt, err := client.NewPoint(
		msg.Id,
		tags,
		fields,
		msg.Time,
	)
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	if err := clnt.Write(bp); err != nil {
		log.Fatal(err)
	}
}

func handleServerConnection(conn net.Conn, c client.Client) {

	for {
		var msg sample
		d := json.NewDecoder(conn)
		err := d.Decode(&msg)
		fmt.Println(msg, err)

		go writePoints(c, msg)
	}
}

func main() {

	fmt.Println("Launching gateway...")

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

	// Create Database
	res, err := queryDB(c, fmt.Sprintf("CREATE DATABASE %s", MyDB))
	if err != nil {
		log.Fatal(res, err)
	}

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
		go handleServerConnection(conn, c)
	}
}
