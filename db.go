package main

import (
	"crypto/tls"
	"net"

	"gopkg.in/mgo.v2"
)

var dbInfo = config.DB.DialInfo()

var dbSession *mgo.Session
var db *mgo.Database

// initDb initialises the dbSession and db variables with a new mongodb session and the default database
func initDb() {
	tlsConfig := &tls.Config{}
	dbInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	dbSession, err := mgo.DialWithInfo(&dbInfo)
	if err != nil {
		panic(err)
	}
	db = dbSession.DB(config.DB.DefaultDB)
}
