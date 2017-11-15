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

	sages := map[string]string{
		"India":      "Ghandi",
		"USA":        "MLK",
		"Meditation": "Buddha",
		"Love":       "Jesus",
		"Prophet":    "Muhammed",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
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

    <li>India - Ghandi</li>

    <li>Love - Jesus</li>

    <li>Meditation - Buddha</li>

    <li>Prophet - Muhammed</li>

    <li>USA - MLK</li>

    </ul>
</body>
</html>
*/
