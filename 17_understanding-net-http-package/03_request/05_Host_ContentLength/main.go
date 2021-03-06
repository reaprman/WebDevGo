package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

var tpl *template.Template

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   url.Values //map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}

/*GET METHOD RUN RESULT:
The Request method: GET

The request host: localhost:8080

The request URL:

Scheme:
Opaque:
Host:
Path: /
RawPath:
ForceQuery: false
RawQuery:
Fragment:
Accept

text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*`/*;q=0.8
Accept-Encoding

gzip, deflate, br
Accept-Language

en-US,en;q=0.9
Cache-Control

max-age=0
Connection

keep-alive
Dnt

1
Upgrade-Insecure-Requests

1
User-Agent

Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36

POST METHOD RUN RESULT:

The Request method: POST

The request host: localhost:8080

The request ContentLength: 33

The request URL:

Scheme:
Opaque:
Host:
Path: /
RawPath:
ForceQuery: false
RawQuery:
Fragment:
Accept

text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*`/*;q=0.8
Accept-Encoding

gzip, deflate, br
Accept-Language

en-US,en;q=0.9
Cache-Control

max-age=0
Connection

keep-alive
Content-Length

33
Content-Type

application/x-www-form-urlencoded
Dnt

1
Origin

http://localhost:8080
Referer

http://localhost:8080/
Upgrade-Insecure-Requests

1
User-Agent

Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36
variable names (identifiers) and values:

fname

RYAN
submit-btn

onda button
*/
