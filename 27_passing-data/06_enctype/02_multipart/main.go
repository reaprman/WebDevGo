package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)

	err := tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

/*
Run Result:

BODY:
------WebKitFormBoundarycl6lUJueleUeAwJL Content-Disposition: form-data; name="first" Master Chief
------WebKitFormBoundarycl6lUJueleUeAwJL Content-Disposition: form-data; name="last" John-117
------WebKitFormBoundarycl6lUJueleUeAwJL Content-Disposition: form-data; name="quote"; filename="keys to success.txt" Content-Type: text/plain The successful warrior is the average man, with laser-like focus. My success, part of it certainly, is that I have focused in on a few things.
------WebKitFormBoundarycl6lUJueleUeAwJL--

*/
