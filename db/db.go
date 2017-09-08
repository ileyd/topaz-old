package db

import (
	"crypto/tls"
	"errors"
	"net"

	"gopkg.in/mgo.v2"
)

//DB ...
type DB struct {
}

var DbInfo = &mgo.DialInfo{
	Addrs: []string{
		"cluster0-shard-00-00-qqlpb.mongodb.net:27017",
		"cluster0-shard-00-01-qqlpb.mongodb.net:27017",
		"cluster0-shard-00-02-qqlpb.mongodb.net:27017",
	},
	Database:       "admin",
	ReplicaSetName: "Cluster0-shard-0",
	Username:       "admin",
	Password:       "VwJnf1lQmASMnVET",
}

var db *mgo.Database
var ErrNotFound = errors.New("not found")

//Init ...
func Init() {
	tlsConfig := &tls.Config{}
	DbInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	session, err := mgo.DialWithInfo(DbInfo)
	if err != nil {
		panic(err)
	}
	db = session.DB("stream")
}

//GetDB ...
func GetDB() *mgo.Database {
	return db
}
