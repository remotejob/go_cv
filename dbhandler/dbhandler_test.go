package dbhandler

import (
	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestInsertRecord(t *testing.T) {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	
	InsertRecord(*session,"en_US","programming","test.com","blog")

}
