//Package config of Thunder
package config

import (
	_ "github.com/lib/pq" //Because
	"gopkg.in/mgo.v2"
	"log"
)

// DB represents the Mongo Database
var DB *mgo.Database

// DBName represents the name of the Database
var DBName = "thunder"

/**
// Collections
var Requests *mgo.Collection
var Responses *mgo.Collection
 */

func init() {
	// get a mongo session
	s, err := mgo.Dial("mongodb://127.0.0.1:27017/" + DBName)
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB(DBName)
	log.Println("[Database] You connected to your mongo database.")
}

// GetCollection function that returns a collection in the database,
// that corresponds to the name inside the database, passed in the string.
func GetCollection(name string) *mgo.Collection {
	return DB.C(name)
}
