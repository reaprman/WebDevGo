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
	body := "Writing Response. Check Response body payload\n"
	body += rMethod
	body += "\n"
	body += rURI
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

/*
Run Result:
GET / HTTP/1.1
***RESPONSE METHOD GET
***RESPONE URL /
Host: localhost:8080
Connection: keep-alive
Cache-Control: max-age=0
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Upgrade-Insecure-Requests: 1
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*`/*;q=0.8
DNT: 1
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

This is END of the HTTP REQUEST HEADERS
*/
