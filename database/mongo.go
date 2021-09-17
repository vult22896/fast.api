package database

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
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
		error := godotenv.Load()
		if error != nil {
			panic("Failed load env file")
		}
		user := os.Getenv("DB_Mongo_User")
		pass := os.Getenv("DB_Mongo_Password")
		host := os.Getenv("DB_Mongo_Host")
		port := os.Getenv("DB_Mongo_Port")
		name := os.Getenv("DB_Mongo_Name")
		instanceMongo = &mongodb{
			DB_Mongo_User:     user,
			DB_Mongo_Password: pass,
			DB_Mongo_Host:     host,
			DB_Mongo_Port:     port,
			DB_Mongo_Name:     name,
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
