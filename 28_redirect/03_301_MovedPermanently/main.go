package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
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
	fmt.Println("This is the request method for bar: ", req.Method)
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("This is the request method for barred: ", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

/*
Run Result:
This is the request method for foo: GET

This is the request method for barred:  GET
This is the request method for bar:  POST
This is the request method for foo: GET

This is the request method for bar:  GET
This is the request method for foo: GET

This is the request method for foo: GET
*/
