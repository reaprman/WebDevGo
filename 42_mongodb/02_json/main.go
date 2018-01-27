package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/reaprman/WEBDEVGO/42_mongodb/02_json/models"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	// added route plus parameter
	r.GET("/user/:id", getUser)
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	s := `<!doctype html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Index</title>
</head>
<body>
<a href="/user9872309847">GO TO: http://localhost:8080/user/9872309847</a>
</body>
</html>`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

// changed fun name
func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	// Marshel into JSON
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

/*
Run Result:
{"name":"James Bond","gender":"male","age":32,"id":"1"}
*/
