package main

import (
	"html/template"
	"net/http"

	"github.com/reaprman/WebDevGo/42_mongodb/10_hands-on/controller"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	sc := controller.NewSessionController(tpl)
	http.HandleFunc("/", sc.Index)
	http.HandleFunc("/bar", sc.Bar)
	http.HandleFunc("/signup", sc.Signup)
	http.HandleFunc("/login", sc.Login)
	http.HandleFunc("/logout", sc.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
