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
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]
	fmt.Println("***METHOD", m)
	fmt.Println("***URI", u)

	// router, multiplexer, etc...
	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/about" {
		about(conn)
	}
	if m == "GET" && u == "/contact" {
		contact(conn)
	}
	if m == "GET" && u == "/apply" {
		apply(conn)
	}
	if m == "POST" && u == "/apply" {
		applyProcess(conn)
	}
}

func index(conn net.Conn) {
	body := `<DOCTYPE html><html lang"en"><head><meta charset="UTF-8">
	<title></title></head><body>
	<strong>INDEX</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<DOCTYPE html><html lang"en"><head><meta charset="UTF-8">
	<title></title></head><body>
	<strong>ABOUT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `<DOCTYPE html><html lang"en"><head><meta charset="UTF-8">
	<title></title></head><body>
	<strong>CONTACT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	body := `<DOCTYPE html><html lang"en"><head><meta charset="UTF-8">
	<title></title></head><body>
	<strong>APPLY</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	<form method="post" action="/apply">
	<input type="submit" value="apply">
	</form>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func applyProcess(conn net.Conn) {
	body := `<DOCTYPE html><html lang"en"><head><meta charset="UTF-8">
	<title></title></head><body>
	<strong>APPLY PROCESS</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

/*
Click through of all links and buttons
Run Result:

ET / HTTP/1.1
***METHOD GET
***URI /
Host: localhost:8080
Connection: keep-alive
Cache-Control: max-age=0
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Upgrade-Insecure-Requests: 1
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*`/*;q=0.8
DNT: 1
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

GET /favicon.ico HTTP/1.1
***METHOD GET
***URI /favicon.ico
Host: localhost:8080
Connection: keep-alive
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Accept: image/webp,image/apng,image/*`,*`/*;q=0.8
DNT: 1
Referer: http://localhost:8080/
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

GET /favicon.ico HTTP/1.1
***METHOD GET
***URI /favicon.ico
Host: localhost:8080
Connection: keep-alive
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Accept: image/webp,image/apng,image/*,*`/*;q=0.8
DNT: 1
Referer: http://localhost:8080/
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

GET /favicon.ico HTTP/1.1
***METHOD GET
***URI /favicon.ico
Host: localhost:8080
Connection: keep-alive
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Accept: image/webp,image/apng,image/*,*`/*;q=0.8
DNT: 1
Referer: http://localhost:8080/
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

GET / HTTP/1.1
***METHOD GET
***URI /
Host: localhost:8080
Connection: keep-alive
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*`/*;q=0.8
DNT: 1
Referer: http://localhost:8080/
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

GET /favicon.ico HTTP/1.1
***METHOD GET
***URI /favicon.ico
Host: localhost:8080
Connection: keep-alive
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Accept: image/webp,image/apng,image/*,*`/*;q=0.8
DNT: 1
Referer: http://localhost:8080/
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

GET /about HTTP/1.1
***METHOD GET
***URI /about
Host: localhost:8080
Connection: keep-alive
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*`/*;q=0.8
DNT: 1
Referer: http://localhost:8080/
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

GET /favicon.ico HTTP/1.1
***METHOD GET
***URI /favicon.ico
Host: localhost:8080
Connection: keep-alive
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Accept: image/webp,image/apng,image/*,*`/*;q=0.8
DNT: 1
Referer: http://localhost:8080/about
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

GET /contact HTTP/1.1
***METHOD GET
***URI /contact
Host: localhost:8080
Connection: keep-alive
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*`/*;q=0.8
DNT: 1
Referer: http://localhost:8080/about
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

GET /favicon.ico HTTP/1.1
***METHOD GET
***URI /favicon.ico
Host: localhost:8080
Connection: keep-alive
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Accept: image/webp,image/apng,image/*,*`/*;q=0.8
DNT: 1
Referer: http://localhost:8080/contact
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

GET /apply HTTP/1.1
***METHOD GET
***URI /apply
Host: localhost:8080
Connection: keep-alive
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*`/*;q=0.8
DNT: 1
Referer: http://localhost:8080/contact
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

GET /favicon.ico HTTP/1.1
***METHOD GET
***URI /favicon.ico
Host: localhost:8080
Connection: keep-alive
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Accept: image/webp,image/apng,image/*,*`/*;q=0.8
DNT: 1
Referer: http://localhost:8080/apply
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

POST /apply HTTP/1.1
***METHOD POST
***URI /apply
Host: localhost:8080
Connection: keep-alive
Content-Length: 0
Cache-Control: max-age=0
Origin: http://localhost:8080
Upgrade-Insecure-Requests: 1
Content-Type: application/x-www-form-urlencoded
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*`/*;q=0.8
DNT: 1
Referer: http://localhost:8080/apply
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

GET /favicon.ico HTTP/1.1
***METHOD GET
***URI /favicon.ico
Host: localhost:8080
Connection: keep-alive
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
Accept: image/webp,image/apng,image/*,*`/*;q=0.8
DNT: 1
Referer: http://localhost:8080/apply
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
*/
