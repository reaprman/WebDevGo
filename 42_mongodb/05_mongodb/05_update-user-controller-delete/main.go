package main

import (
	"net/http"

	"github.com/reaprman/WebDevGo/42_mongodb/05_mongodb/05_update-user-controller-delete/controllers"

	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/", index)
	// added route
	r.POST("/user", uc.CreateUser)
	r.GET("/user/:id", uc.GetUser)
	r.DELETE("/user/:id", uc.DeleteUser)
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

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

// changed fun name

/*
Run Result:
{"name":"James Bond","gender":"male","age":32,"id":"007"}

Write code to delete user

*/
