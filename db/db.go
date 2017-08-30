package db

import (
	"gopkg.in/mgo.v2"
)

//DB ...
type DB struct {
}

const (
	//DbUser ...
	DbUser = "topaz"
	//DbPassword ...
	DbPassword = "respiteforever"
	//DbName ...
	DbName = "stream"
	//DbHost ...
	DbHost = "cluster0-shard-00-00-a9ipr.mongodb.net:27017,cluster0-shard-00-01-a9ipr.mongodb.net:27017,cluster0-shard-00-02-a9ipr.mongodb.net:27017"
	//DbURL ...
	DbURL = "mongodb://" + DbUser + ":" + DbPassword + "@" + DbHost + "/" + DbName + "/?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin"
)

var db *mgo.Database

//Init ...
func Init() {
	var err error
	db, err = ConnectDB(DbURL)
	if err != nil {
		panic(err)
	}
}

//ConnectDB ...
func ConnectDB(url string) (*mgo.Database, error) {
	session, error := mgo.Dial(url)
	return session.DB(DbName), error
}

//GetDB ...
func GetDB() *mgo.Database {
	return db
}
