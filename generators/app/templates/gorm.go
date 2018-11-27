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
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DbUsername, config.DbPassword, config.DbHostname, config.DbPort, config.DbName)

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

	return DB, err
}

/*
   NOTE: test DB will be dropped each time OpenTestDB is called
*/

//OpenTestDB : Opens a new connection and creates test database
func OpenTestDB() (*gorm.DB, error) {

	//Set default test db name
	testDBName := config.TestDBName
	if testDBName == "" {
		testDBName = "test"
	}

	//Create the Url used to Open the db
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", "root", config.DbRootPassword, config.DbHostname, config.DbPort)

	//Attempt to open a new connect to the db
	glog.Info("Opening a connection to the db...")
	db, err := gorm.Open(config.DbDriver, dbURL)
	if err != nil {
		glog.Fatal("Couldn't open a connection to the db!", err)
		return nil, err
	}

	//Drop old test db
	dropDBStmt := fmt.Sprintf(`DROP DATABASE IF EXISTS %s`, config.TestDBName)
	db.Exec(dropDBStmt)

	//Create Database
	createStmt := fmt.Sprintf(`CREATE DATABASE %s;`, config.TestDBName)
	result := db.Exec(createStmt)
	if result.Error != nil {
		glog.Info(result.Error)
		return nil, err
	}

	//Select new test database
	useStmnt := fmt.Sprintf(`USE %s;`, config.TestDBName)
	result = db.Exec(useStmnt)
	if result.Error != nil {
		glog.Info(result.Error)
		return nil, err
	}

	db.LogMode(true)

	TestDB = db

	return db, nil
}
