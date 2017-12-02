package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "I dialed you.")

}

/*
Run with 02_read to produce the following output,
ran this program 3 times

Run Result:
I dialed you.
We Got Here
I dialed you.
We Got Here
I dialed you.
We Got Here
*/
