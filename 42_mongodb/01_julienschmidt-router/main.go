package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

/*
Browser
Run Results:
Welcome!
*/
