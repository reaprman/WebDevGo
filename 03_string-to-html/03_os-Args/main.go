package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello Wordl!</title>
		</head>
		<body>
		<h1> ` +
		name +
		`</h1
		</body>
		</html>
		`)
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error creating file", err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(str))
}

/*
at terminal to use os.Args:
go run main.go Ryan
Run Result:
C:\Users\ryang\AppData\Local\Temp\go-build614672274\command-line-arguments\_obj\exe\main.exe
Ryan

also new html file created index.html that contins html code in this program
*/
