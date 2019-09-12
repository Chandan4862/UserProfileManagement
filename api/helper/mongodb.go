package handler

import(
	"log"
	"gopkg.in/mgo.v2"
)
var db *mgo.Database

func InitDb(){

session,err := mgo.Dial("localhost")
if err!=nil{
	log.Fatalf("Failed connection %v", err)
}
db= session.DB("user_profile")
}

func Collection() *mgo.Collection{
	return db.C("users")
}