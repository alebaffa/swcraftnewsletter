package database

import "gopkg.in/mgo.v2"

func ciao() {
	session, _ := mgo.Dial("asdasdad")
    defer session.Close
    c := session.DB("")
}
