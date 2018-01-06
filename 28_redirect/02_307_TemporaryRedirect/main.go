package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("This is the request method for foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("The request method for bar: ", req.Method)
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("The request method for barred: ", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

/*
Temporary redirect preserves request method that it is orginally sent with.

Run Result:
This is the request method for foo: GET

The request method for barred:  GET
The request method for bar:  POST
This is the request method for foo: POST

The request method for bar:  GET
This is the request method for foo: GET
*/
