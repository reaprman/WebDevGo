package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Name   string        `json:"name"`
	Gender string        `json:"gender"`
	Age    int           `json:"age"`
	Id     bson.ObjectId `json:"id" bson:"_id"`
}

// ID was of type string before
