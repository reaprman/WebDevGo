package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Logan-Key", "This is from logan")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any Code you want in this func to respond with</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}

/*
Run Result:

Any Code you want in this func to respond with
*/
