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

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King Jr.",
		Motto: "Hatred never ceases with hatred but with love alone is healed",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	muhammed := sage{
		Name:  "Muhammed",
		Motto: "",
	}

	sages := []sage{buddha, gandhi, mlk, jesus, muhammed}

	err := tpl.Execute(os.Stdout, sages)
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
    <title>My Peeps</title>
</head>
<body>
    <ul>

    <li>0: Buddha - The belief of no beliefs</li>

    <li>1: Gandhi - Be the change</li>

    <li>2: Martin Luther King Jr. - Hatred never ceases with hatred but with love alone is hea
led</li>

    <li>3: Jesus - Love all</li>

    <li>4: Muhammed - </li>

    </ul>
</body>
</html>
*/
