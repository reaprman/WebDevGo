package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	//using must to avoid checking err. Must does this for us
	//review Must documentation
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Release self-focus: embrace other-focus.`)
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
        <title>Hello World</title>
    </head>
    <body>

        <h1>The meaning of life: Release self-focus: embrace other-focus.</h1?
    </body>
</html>
*/
