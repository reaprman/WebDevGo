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

	sages := []string{"Gandhi", "MLK", "Jesus", "Buddha", "Muhammed"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}

/*
This can be written into its on html file using os.Create(

)
Run Result:
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Peeps</title>
</head>
<body>
    <ul>

    <li>0 - Gandhi</li>

    <li>1 - MLK</li>

    <li>2 - Jesus</li>

    <li>3 - Buddha</li>

    <li>4 - Muhammed</li>

    </ul>
</body>
</html>
*/
