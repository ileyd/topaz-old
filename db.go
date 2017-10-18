package main

import (
	"crypto/tls"
	"log"
	"net"
	"time"

	"gopkg.in/mgo.v2"
)

var dbSession *mgo.Session
var db *mgo.Database

// initDb initialises the dbSession and db variables with a new mongodb session and the default database
func initDb() (err error) {
	var dbInfo = &mgo.DialInfo{
		Addrs:          config.DB.Addrs,
		ReplicaSetName: config.DB.ReplicaSetName,
		Database:       config.DB.AuthDB,
		Username:       config.DB.Username,
		Password:       config.DB.Password,
		Timeout:        time.Second * 5,
	}

	dbInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		tlsConfig := &tls.Config{}
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	dbSession, err = mgo.DialWithInfo(dbInfo)
	if err == nil {
		log.Println("successfully connected to the database")
	}
	db = dbSession.DB(config.DB.AppDB)
	return err
}
