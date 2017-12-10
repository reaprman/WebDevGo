package main

import (
	"io"
	"net/http"
)

func def(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Default entry point")
}

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

func m(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, Ryan Logan!")
}

func main() {

	http.HandleFunc("/", def)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", m)

	http.ListenAndServe(":8080", nil)
}

/*
Run Result:

*/
