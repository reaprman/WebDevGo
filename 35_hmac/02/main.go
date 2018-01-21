package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/authenticate", auth)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-session")
	if err != nil {
		c = &http.Cookie{
			Name:  "my-session",
			Value: "",
		}
	}

	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		c.Value = e + `|` + getCode(e)
	}
	http.SetCookie(w, c)

	io.WriteString(w, `<!doctype html>
		<html>
		<body>
		<form method="post">
			<input type="email" name="email">
			<input type="submit">
		</form>
		<a href="/authenticate">Validate This `+c.Value+`</a>
		</body>
		</html>`)
}

func auth(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if c.Value == "" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	xs := strings.Split(c.Value, "|")
	email := xs[0]
	codeRcvd := xs[1]
	codeCheck := getCode(email + "s")

	if codeRcvd != codeCheck {
		fmt.Println("HMAC codes didn't match")
		fmt.Println(codeRcvd)
		fmt.Println(codeCheck)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	io.WriteString(w, `<!doctype html>
		<html>
		  <body>
			<h1>`+codeRcvd+` - RECEIVED </h1>
			<h1>`+codeCheck+` - RECALCULATED </h1>
		  </body>
		</html>`)
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

/*
HMAC WILL NEVER MATCH AS EXTRA CHAR (s) is added to the email

Run Result:
HMAC codes didn't match
453159a62580804892cc90a27dfb8bfdf2309107336445dcfd0186674111ee71
826706a18972415ed6d49207a040a55bb17cb7a7627093f23d7148c14b42f2e0

FROM BROWSER AFTER  correcting above:
453159a62580804892cc90a27dfb8bfdf2309107336445dcfd0186674111ee71 - RECEIVED
453159a62580804892cc90a27dfb8bfdf2309107336445dcfd0186674111ee71 - RECALCULATED
*/
