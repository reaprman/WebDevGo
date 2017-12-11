package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	var i int
	var rMethod, rURI string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// First line is REQUEST LINE
			rMethod = strings.Fields(ln)[0]
			rURI = strings.Fields(ln)[1]
			fmt.Println("***RESPONSE METHOD", rMethod)
			fmt.Println("***RESPONE URL", rURI)
		}
		if ln == "" {
			// end of headers
			fmt.Println("This is END of the HTTP REQUEST HEADERS")
			break
		}
		i++
	}
	body := "<h1>HOLY COW THIS IS LOW LEVEL</h1>"
	body += rMethod
	body += "<br>"
	body += rURI
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

/*
Run Result:
HOLY COW THIS IS LOW LEVEL

GET
/
*/
