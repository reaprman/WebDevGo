package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	// now we get here
	// the connection will timeout
	// that breaks us out of the scanner loop
	fmt.Println("***CODE GOT HERE***")
}

/*
The scanner loop within the handler func times out after 10 seconds

From CMD Run Result:
I heard you say: There is no cursor the first time


Connection to host lost.

From Console Run Result:
There is no cursor the first time
***CODE GOT HERE***
*/
