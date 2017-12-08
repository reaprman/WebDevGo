package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog doggy dogggy")
}

func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat kitty kittty")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()

	// trailing / on dog allows it to handle /dog/something/else
	// /cat will not handle anything other than /cat
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}

/*
Run Result from /dog:
dog doggy dogggy

Run Result from /dog/also/handles/this:
dog doggy dogggy

Run Result from /cat:
cat kitty kittty

Run Result from /cat/does/not/handle/this:
404 page not found

*/
