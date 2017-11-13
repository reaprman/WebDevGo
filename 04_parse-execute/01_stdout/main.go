package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
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
                <title>Hello Wordl!</title>
        </head>
        <body>
                <h1>Hello Ryan</h1
        </body>
</html>
*/
