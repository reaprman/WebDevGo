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
	// added route
	r.POST("/user", createUser)
	// added route plus parameter
	r.GET("/user/:id", getUser)
	r.DELETE("/user/:id", deleteUser)
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
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// composite literal - type and curly brace
	u := models.User{}

	// encode/decode for sending/receiving JSON to/from a stream
	json.NewDecoder(r.Body).Decode(&u)

	// change id
	u.Id = "007"

	// marshall/unmarshal for having JSON assigned to a variable
	uj, _ := json.Marshal(u)

	// write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201ssss
	fmt.Fprintf(w, "%s\n", uj)
}

func deleteUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// TODO: write code to delete user
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user")

}

/*
Run Result:
{"name":"James Bond","gender":"male","age":32,"id":"007"}

Write code to delete user

*/
