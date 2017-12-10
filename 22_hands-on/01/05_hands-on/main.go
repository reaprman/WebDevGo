package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func def(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Default entry point")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("index.gohtml"))

	err := tpl.ExecuteTemplate(w, "index.gohtml", "Ryan")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(def))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}

/*
Run Result:

Default entry point

dog dog doggy

Hello Ryan Welcome once again Ryan
*/
