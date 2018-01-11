package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("visitCount")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "visitCount",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++

	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w, cookie)

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
	fmt.Fprintln(w, "COOKIE VALUE: ", cookie.Value)
}

/*
Result from web browser

Run Result:

COOKIE WRITTEN - CHECK YOUR BROWSER
in chrome go to: dev tools / application / cookies
COOKIE VALUE:  1

COOKIE WRITTEN - CHECK YOUR BROWSER
in chrome go to: dev tools / application / cookies
COOKIE VALUE:  2
*/
