package database

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Sql interface {
	Connect() *gorm.DB
}
type sql struct {
	DB_User     string
	DB_Password string
	DB_Name     string
	DB_Host     string
}

var instanceMysql *sql
var onceMySql sync.Once

func GetInstanceMysql() Sql {
	onceMySql.Do(func() {
		error := godotenv.Load()
		if error != nil {
			panic("Failed load env file")
		}

		dbUser := os.Getenv("DB_USERNAME")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")

		instanceMysql = &sql{
			DB_User:     dbUser,
			DB_Password: dbPassword,
			DB_Name:     dbName,
			DB_Host:     dbHost,
		}
	})
	return instanceMysql
}

func (dbsql *sql) Connect() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dbsql.DB_User + ":" + dbsql.DB_Password + "@tcp(" + dbsql.DB_Host + ")/" + dbsql.DB_Name + "?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                                                                                 // default size for string fields
		DisableDatetimePrecision:  true,                                                                                                                                // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                                                                                // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                                                                                // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                                                                               // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Connect mysql error")
	} else {
		fmt.Println("connect success")
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Connect poll mysql error")
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
