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
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// First line is REQUEST LINE
			mux(conn, ln)
		}
		if ln == "" {
			// end of headers
			fmt.Println("This is END of the HTTP REQUEST HEADERS")
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	rMethod := strings.Fields(ln)[0]
	rURI := strings.Fields(ln)[1]
	fmt.Println("***RESPONSE METHOD", rMethod)
	fmt.Println("***RESPONSE URI", rURI)
	res := []string{rMethod, rURI}

	if rMethod == "GET" && rURI == "/" {
		def(conn, res)
	}

	if rMethod == "GET" && rURI == "/apply" {
		apply(conn, res)
	}

	if rMethod == "POST" && rURI == "/apply" {
		applyProcess(conn, res)
	}

}

func def(conn net.Conn, res []string) {
	body := `<DOCTYPE html>
	<html lang"en">
		<head><meta charset="UTF-8">
		<title></title>
	</head>
	<body>
		<strong>DEFAULT</strong><br>
		<a href="/">default</a><br>
		<a href="/apply">apply</a><br>
		<p>`
	body += res[0]
	body += "<br>"
	body += res[1]
	body += "</p></body></html>"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func apply(conn net.Conn, res []string) {
	body := `<DOCTYPE html>
	<html lang"en">
	<head>
		<meta charset="UTF-8">
		<title></title>
	</head>
	<body>
		<strong>APPLY</strong><br>
		<a href="/">default</a><br>
		<a href="/apply">apply</a><br>
		<form method="post" action="/apply">
		<input type="submit" value="apply">
		</form>
		<p>`
	body += res[0]
	body += "<br>"
	body += res[1]
	body += "</p></body></html>"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func applyProcess(conn net.Conn, res []string) {
	body := `<DOCTYPE html>
	<html lang"en">
	<head>
		<meta charset="UTF-8">
		<title></title>
	</head>
	<body>
		<strong>APPLY</strong><br>
		<a href="/">default</a><br>
		<a href="/apply">apply</a><br>
		<p>`
	body += res[0]
	body += "<br>"
	body += res[1]
	body += "</p></body></html>"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
