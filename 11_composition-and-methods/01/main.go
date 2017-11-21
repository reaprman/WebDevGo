package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1 := person{
		Name: "James Bond",
		Age:  42,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}

/*
Run Result:

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Compisition</title>
</head>
<body>

<h1>James Bond</h1>
<h2>42</h2
</body>
</html>
*/
