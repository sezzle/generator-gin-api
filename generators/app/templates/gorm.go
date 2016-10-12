package gorm

import (
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	// recommended by gorm to have this blank import
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func init() {
	SetDBSettingsFromViper()
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
	//DbDriver : The Driver for the DB - Mysql for now, although we can test others.package settings
	DbDriver string
	// DbUsername : Username to loging to db
	DbUsername string
	// DbPassword : Password for db
	DbPassword string
	//DbRootPassword :  Root Password for db, used to create test db *Local Env Only
	DbRootPassword string
	//DbHostname :  Location of the db in the cluster. Usually surrounded by tcp()
	DbHostname string
	//DbPort : Port that our db is open at - traditionally 3306
	DbPort string
	//DbName : The name of the specific db.
	DbName string
	// TestDBName : Default: test | Set in local environment to avoid name clash
	TestDBName string
	//Debug : Puts database in debug mode
	Debug string
	//DBSettings is an object to reference for database configuration
	DBSettings = &dbConfig{}
)

//DbConfig : A struct that holds all the configuration for the db - from viper variables. //CURRENTLY DEPRECATED
type dbConfig struct {
	DbDriver string //The Driver for the DB - Mysql for now, although we can test others.package settings

	DbUsername     string        //Username to loging to db
	DbPassword     string        //Password for db
	DbRootPassword string        //Root Password for db, used to create test db *Local Env Only
	DbHostname     string        //Location of the db in the cluster. Usually surrounded by tcp()
	DbPort         string        //Port that our db is open at - traditionally 3306
	DbName         string        //THe name of the specific db.
	numOfReq       int           //Number of Requests to ping to the db while waiting
	timePerReq     time.Duration //Amount of time per request above. numOfReq * timePerReq = total time waiting for db to connect before throwing an error
	TestDBName     string        //Default: test | Set in local environment to avoid name clash
	Debug          string        //Puts database in debug mode
}

//InitDB : This Initalizes the first db, and exports it to be passed around.
func InitDB() (*gorm.DB, error) {

	//Create the Url used to Open the db
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DbUsername, DbPassword, DbHostname, DbPort, DbName)

	//Attempt to open a new connect to the db
	glog.Info("Opening a connection to the db...")
	db, err := gorm.Open(DbDriver, dbURL)
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
	testDBName := TestDBName
	if testDBName == "" {
		testDBName = "test"
	}

	//Create the Url used to Open the db
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", "root", DbRootPassword, DbHostname, DbPort)

	//Attempt to open a new connect to the db
	glog.Info("Opening a connection to the db...")
	db, err := gorm.Open(DbDriver, dbURL)
	if err != nil {
		glog.Fatal("Couldn't open a connection to the db!", err)
		return nil, err
	}

	//Drop old test db
	dropDBStmt := fmt.Sprintf(`DROP DATABASE IF EXISTS %s`, TestDBName)
	db.Exec(dropDBStmt)

	//Create Database
	createStmt := fmt.Sprintf(`CREATE DATABASE %s;`, TestDBName)
	result := db.Exec(createStmt)
	if result.Error != nil {
		glog.Info(result.Error)
		return nil, err
	}

	//Select new test database
	useStmnt := fmt.Sprintf(`USE %s;`, TestDBName)
	result = db.Exec(useStmnt)
	if result.Error != nil {
		glog.Info(result.Error)
		return nil, err
	}

	db.LogMode(true)

	TestDB = db

	return db, nil
}

//SetDBSettingsFromViper : A Function that sets the global settings structs
//using the viper configuration package
func SetDBSettingsFromViper() {
	glog.Info("Generating database settings...")

	if os.Getenv("SEZZLE_ENVIRONMENT") == "docker" {
		//Try docker settings
		viper.SetConfigName("dockerConfig")
		viper.AddConfigPath("/go/src/sezzle/instantach/config/")
		viper.SetConfigType("yaml")

		err := viper.ReadInConfig()
		if err != nil {
			glog.Fatal("Could not properly load docker settings: ", err)
		}
	} else if os.Getenv("SEZZLE_ENVIRONMENT") == "staging" {
		glog.Info("Reading from staging settings")

		DbDriver = os.Getenv("DB_DRIVER")
		DbHostname = os.Getenv("DB_HOSTNAME")
		DbUsername = os.Getenv("DB_USERNAME")
		DbPort = os.Getenv("DB_PORT")
		DbName = os.Getenv("DB_NAME")
		DbPassword = os.Getenv("DB_PASSWORD")
		Debug = os.Getenv("DEBUG")

		// viper.SetConfigName("stagingConfig")
		// viper.AddConfigPath("./config/")
		// viper.SetConfigType("yaml")

		// err := viper.ReadInConfig()
		// if err != nil {
		// 	glog.Info("Failed reading local settings: ", err)
		// }
		return

	} else {
		//Try local settings first
		glog.Info("Reading from local settings")
		viper.SetConfigName("localConfig")
		viper.AddConfigPath("./config/")
		viper.SetConfigType("yaml")

		err := viper.ReadInConfig()
		if err != nil {
			glog.Info("Failed reading local settings: ", err)
		}
	}

	//Set db config object
	DbUsername = viper.GetString("dbUsername")
	DbPassword = viper.GetString("dbPassword")
	DbRootPassword = viper.GetString("dbRootPassword")
	DbHostname = viper.GetString("dbHostName")
	DbName = viper.GetString("dbName")
	DbPort = viper.GetString("dbPort")
	DbDriver = viper.GetString("dbDriver")
	TestDBName = viper.GetString("testDbName")
	Debug = viper.GetString("debug")

}
