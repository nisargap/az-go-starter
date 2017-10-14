package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id             bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username       string        `json:"username" bson:"email"`
	Salt           string        `json:"salt,omitempty" bson:"salt,omitempty"`
	Privilege      string        `json:"privilege" bson:"privilege"`
	Password       string        `json:"password,omitempty" bson:"password,omitempty"`
	ProfileImage   string        `json:"profile_img" bson:"profile_img"`
	DateRegistered string        `json:"date_registered" bson:"date_registered"`
}
