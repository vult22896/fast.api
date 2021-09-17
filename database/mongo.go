package database

import (
	"fmt"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
)

type MongoDB interface {
	Connect() *mgo.Session
}

type mongodb struct {
	DB_Mongo_User     string
	DB_Mongo_Password string
	DB_Mongo_Host     string
	DB_Mongo_Port     string
	DB_Mongo_Name     string
}

var (
	instanceMongo *mongodb
	onceMongo     sync.Once
)

func GetInstanceMongo() MongoDB {
	onceMongo.Do(func() {
		instanceMongo = &mongodb{
			DB_Mongo_User:     "",
			DB_Mongo_Password: "",
			DB_Mongo_Host:     "127.0.0.1",
			DB_Mongo_Port:     "27017",
			DB_Mongo_Name:     "bibabo",
		}
	})
	return instanceMongo
}

func (dbmongo *mongodb) Connect() *mgo.Session {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{dbmongo.DB_Mongo_Host},
		Username: dbmongo.DB_Mongo_User,
		Password: dbmongo.DB_Mongo_Password,
		Timeout:  60 * time.Second,
	})

	if err != nil {
		fmt.Printf("[ConnectDB]: %s\n", err)
	}
	session.SetMode(mgo.Monotonic, true)

	return session
}
