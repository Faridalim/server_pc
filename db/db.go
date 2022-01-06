package db

import (
	"database/sql"
	"fmt"
	"server_pc/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	// username:password@protocol(address)/dbname?param=value
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?parseTime=true&loc=Asia%2FJakarta"
	fmt.Println(connectionString)

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("connection string error")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
