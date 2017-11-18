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

	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"Logan",
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}

/*
Run Result:
<!DOCTYPE>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Predefined Global Functions</title>
</head>
<body>

    zero
    one
    two

    <!--
    index
    a func you can ue in a template
    it is a "predefined global function"

    Index returns the result of indexing its first agument by the
    following arguments. Thus "index x 123" is, in Go syntax,
    x[1][2][3]. Ech indexed item must be a map, slice, or array.

    https://godoc.org/text/template#hdr-Functions
    -->

    <!-- FYI -->
    This is a go template comment


</body>
</html>
*/
