package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

/* Create a FuncMap to register functions.
"uc" is what the func will be called in the template
"uc" is the ToUpper func from package strings
"ft" is a func we declare
"ft" slices a string, returning the first three characters*/
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

// funcs uc and ft must be added to tpl before you parse the files
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
	// doesn't work tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	//doesn't work tpl = tpl.Funcs(fm)
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {

	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King Jr.",
		Motto: "Hatred never ceases with hatred but with love alone is healed",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
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
    <title>Functions</title>
</head>
<body>

BUDDHA

GANDHI

MARTIN LUTHER KING JR.



Bud

Gan

Mar


</body>
</html>
*/
