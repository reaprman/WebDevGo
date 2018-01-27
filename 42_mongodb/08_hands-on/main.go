package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/reaprman/WebDevGo/42_mongodb/08_hands-on/controllers"
	"github.com/reaprman/WebDevGo/42_mongodb/08_hands-on/models"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[bson.ObjectId]models.User {
	// Create map to hold data
	return models.LoadUser()
}
