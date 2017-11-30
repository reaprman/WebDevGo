package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	// we never ger here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here")
}

/*
From CMD Run Result:
I heard you say: hellohehellodfheljl
                                    How are ou
I heard you say: How are ou
                           I'm doing ok but could be better
I heard you say: I'm doing ok but could be better
                                                 strange
I heard you say: strange

From Console Run Result:
hellohehellodfheljl
How are ou
I'm doing ok but could be better
strange

*/
