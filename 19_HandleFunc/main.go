package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog doggy dogggy")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat kitty kittty")
}

func main() {

	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}

/*
The above code is the simplest version when using a mux.
Compare to 02_NewServeMux and 01 in 03_DefaultServeMux

Run Result from /dog:
dog doggy dogggy

Run Result from /dog/also/handles/this:
dog doggy dogggy

Run Result from /cat:
cat kitty kittty

Run Result from /cat/does/not/handle/this:
404 page not found

*/
