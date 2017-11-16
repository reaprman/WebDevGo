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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no belief",
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

	err := tpl.Execute(os.Stdout, data)
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
    <title>Wisdom & Transport</title>
</head>
<body>
    <ul>

        <li>Buddha - The belief of no belief</li>

        <li>Gandhi - Be the change</li>

        <li>Martin Luther King Jr. - Hatred never ceases with hatred but with love alone is he
aled</li>

    </ul>

    <ul>

        <li>Ford - F150 - 2

        <li>Toyota - Corolla - 4

    </ul>
</body>
</html>
*/
