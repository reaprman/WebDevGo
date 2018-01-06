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
	fmt.Print("Your request method type for foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method type for bar: ", req.Method)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method type for barred: ", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

/*
Run Result:

Your request method type for foo: GET

Your request method type for bar:  GET
Your request method type for foo: GET

Your request method type for barred:  GET
Your request method type for bar:  POST
Your request method type for foo: GET
*/
