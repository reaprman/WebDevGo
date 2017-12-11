package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			if ln == "" {
				fmt.Println("This is END of the HTTP REQUEST HEADERS")
				break
			}
			fmt.Println(ln)
		}
		//defer conn.Close()
		fmt.Println("We got here")
		io.WriteString(conn, "I see you connected")
		conn.Close()
	}
}

/*
io.WriteString did not write to web browser?
Chrome WEb browser says localhost sent invalid ERR_INVALID_HTTP_RESPONSE
This displays with no issues in mozilla Firefox!

Run Result:
GET / HTTP/1.1
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
We got here
*/
