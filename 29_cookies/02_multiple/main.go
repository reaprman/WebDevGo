package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/multiple", multiple)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-goal",
		Value: "front-end-web-dev",
	})

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {

	c1, err := req.Cookie("my-goal")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIE #1: ", c1)

	c2, err := req.Cookie("general")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIE #2: ", c2)

	c3, err := req.Cookie("specific")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIE #3: ", c3)
}

func multiple(w http.ResponseWriter, req *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "favorite thing: technology",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "future job: web dev",
	})

	fmt.Fprintln(w, "MULTIPLE COOKIES WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

/*
Results are from browser

Run Results:

COOKIE WRITTEN - CHECK YOUR BROWSER
in chrome go to: dev tools / application / cookies

MULTIPLE COOKIES WRITTEN - CHECK YOUR BROWSER
in chrome go to: dev tools / application / cookies

YOUR COOKIE #1:  my-goal=front-end-web-dev
YOUR COOKIE #2:  general=favorite thing: technology
YOUR COOKIE #3:  specific=future job: web dev
*/
