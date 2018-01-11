package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	fmt.Fprint(w, `<h1><a href="/set">set a cookie</a></h1>`)
}

func set(w http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("my-goal")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-goal",
			Value: "Front End Web Dev",
		}
	}
	http.SetCookie(w, cookie)

	fmt.Fprintln(w, `<h1><a href="/read">read</a></h1>`)
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-goal")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, `<h1>Your Cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, c)
}

func expire(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-goal")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

/*

Run Results:

Your Cookie:
my-goal=front-end-web-dev
*/
