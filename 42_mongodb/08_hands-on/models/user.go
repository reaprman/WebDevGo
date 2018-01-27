package models

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}

func StoreUser(m map[bson.ObjectId]User) {
	f, err := os.Create("db-file")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	json.NewEncoder(f).Encode(m)
}

func LoadUser() map[bson.ObjectId]User {

	m := make(map[bson.ObjectId]User)

	f, err := os.Open("db-file")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		fmt.Println(err, " here")
	}
	return m
}

// Id was of type string before
