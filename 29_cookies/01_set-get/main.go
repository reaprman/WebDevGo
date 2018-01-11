package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-goal",
		Value: "Front-end-web-dev",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-goal")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIE: ", c)
}

/*
Results from webpage

Run Result:

COOKIE WRITTEN - CHECK YOUR BROWSER
in chrome go to: dev tools / application / cookies

YOUR COOKIE:  my-goal=Front-end-web-dev
*/
