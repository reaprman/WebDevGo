package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}

/*
Run Results:
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Peeps</title>
</head>
<body>
    <ul>


    <li>Buddha - The belief of no beliefs</li>
    </ul>
</body>
</html>
*/
