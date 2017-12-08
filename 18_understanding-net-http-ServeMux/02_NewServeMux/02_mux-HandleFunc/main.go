package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog doggy dogggy")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat kitty kittty")
}

func main() {

	mux := http.NewServeMux()

	// trailing / on dog allows it to handle /dog/something/else
	// /cat will not handle anything other than /cat
	mux.HandleFunc("/dog/", d)
	mux.HandleFunc("/cat", c)

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
