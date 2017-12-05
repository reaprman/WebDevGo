package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this function")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}

/*
Below text served from webpage when going to localhost:8080
Run Result:

Any code you want in this function
*/
