package main

import "fmt"

func main() {
	name := "Ryan Logan"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello Wordl!</title>
	</head>
	<body>
	<h1> ` + name + `</h1
	</body>
	</html>
	`
	fmt.Println(tpl)
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
        <h1> Ryan Logan</h1
        </body>
        </html>
*/
