package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	xs := []string{"zero", "one", "two", "three", "four", "five"}

	err := tpl.Execute(os.Stdout, xs)
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
    <title>Predifined global functions</title>
</head>
<body>
    two
    zero
    one

    <!--
    index
    a func you can use in a template
    it is a "predifined global function"

    Index returns the result of indexing its first argument by the
    following arguments. Thus "index X 1 2 3" is, in Go syntax,
    x[1][2][3]. Each indexed item must be a map, slice, or array.

    https://godoc.org/text/template#hdr-functions
    -->

    <!-- FYI -->
    This is a go template comment


</body>
</html>

*/