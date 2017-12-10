package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func def(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Default entry point")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

func m(w http.ResponseWriter, req *http.Request) {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
	err := tpl.ExecuteTemplate(w, "index.gohtml", "Ryan!")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", def)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", m)
	http.ListenAndServe(":8080", nil)
}

/*
Run Result:
Default entry point
dog dog doggy
Hello Ryan! Welcome
*/
