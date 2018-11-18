package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/reaprman/WebDevGo/42_mongodb/08_hands-on/models"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	// session *mgo.Session
	database map[bson.ObjectId]models.User
}

func NewUserController(db map[bson.ObjectId]models.User) *UserController {
	return &UserController{db}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId hex representation, otherwise return status not found
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// ObjectIdHex returns an ObjectId from the provided hex representation.
	oid := bson.ObjectIdHex(id)

	// composite literal
	u := models.User{}

	// Fetch user
	//if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
	//	w.WriteHeader(404)
	//	return
	//}
	u, ok := uc.database[oid]
	if !ok {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create bson ID
	u.Id = bson.NewObjectId()

	// store the user in mongodb
	//uc.session.DB("go-web-dev-db").C("users").Insert(u)
	uc.database[u.Id] = u
	models.StoreUser(uc.database)

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	// Delete user
	// if err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oid); err != nil {
	//	w.WriteHeader(404)
	//	return
	//}

	_, ok := uc.database[oid]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(uc.database, oid)
	models.StoreUser(uc.database)
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}