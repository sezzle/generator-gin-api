package gorm

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
  "<%= myAppPath %>/config"

	// recommended by gorm to have this blank import
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func init() {
}

//Constants defined across
const (
	numOfReq   int           = 25 //The number of requests to ping the db while waiting for startup
	timePerReq time.Duration = 5  //The amount of time to wait per request. num*time = total time to wait for db initalize.
)

var (
	// DB is connectino handle for the db
	DB *gorm.DB
	//TestDB is connection handle for a test database
	TestDB *gorm.DB
)

//InitDB : This Initalizes the first db, and exports it to be passed around.
func InitDB() (*gorm.DB, error) {

	//Create the Url used to Open the db
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=UTC", config.DbUsername, config.DbPassword, config.DbHostname, config.DbPort, config.DbName)

	//Attempt to open a new connect to the db
	glog.Info("Opening a connection to the db...")
	db, err := gorm.Open(config.DbDriver, dbURL)
	if err != nil {
		glog.Info("Couldn't open a connection to the db!", err)
		return nil, err
	}

	// GOING TO KEEP IN DEBUG MODE REGARDLESS OF DEBUG SETTING
	//Debug settings for db
	// if Debug == "true" {
	db.LogMode(true)
	// }

	//Set our Variable to use this connection
  DB = db

  // Limit our idle connections to 10, with a maximum lifetime of 30 seconds
	db.DB().SetMaxIdleConns(100)
	db.DB().SetConnMaxLifetime(30 * time.Second)

	// Limit total connections to 10
	db.DB().SetMaxOpenConns(100)

	return DB, err
}
