package main

import (
	"bufio"
	"fmt"
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
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write response
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0]
			u := strings.Fields(ln)[1]
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", u)
		}
		if ln == "" {
			// header is done, break from for loop
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8">
	<title></title></head><body><strong>Hello World</strong></body></html>`
	fmt.Fprint(conn, "HTTP1/1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, body)
}

/*
Run Result:

GET / HTTP/1.1
Host: localhost:8080
***URL localhost:8080
Connection: keep-alive
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Upgrade-Insecure-Requests: 1
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*`/*;q=0.8
DNT: 1
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

GET /favicon.ico HTTP/1.1
Host: localhost:8080
***URL localhost:8080
Connection: keep-alive
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Accept: image/webp,image/apng,image/*,*`/*;q=0.8
DNT: 1
Referer: http://localhost:8080/
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
*/
