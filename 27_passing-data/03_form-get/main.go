package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<form method="get">
		<input type="text" name="q">
		<input type="submit">
		</form><br>`+v)
}

/*
Run Result:
http://localhost:8080/?q=john+117
john 117
*/