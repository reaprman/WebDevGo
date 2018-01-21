package main

import (
	"html/template"
	"net/http"

	"github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "my-cookie",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

/*
Run Result:
COOKIE
40d7fe3b-3551-4e7b-b1d8-0aca211b067e
*/
