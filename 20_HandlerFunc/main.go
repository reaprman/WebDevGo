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

	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}

/*
http.HandlerFunc allows conversion of func d and c to
a Handler() type which is what the second parameter
of http.Handle takes

func type -> Handler Type to be used with http.Handle


Run Result from /dog:
dog doggy dogggy

Run Result from /dog/also/handles/this:
dog doggy dogggy

Run Result from /cat:
cat kitty kittty

Run Result from /cat/does/not/handle/this:
404 page not found
*/
