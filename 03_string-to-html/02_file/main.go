package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Ryan Logan"

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
Run Result:
New file created index.html tht contains html code
*/
