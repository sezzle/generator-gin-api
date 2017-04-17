package config

import (
  "fmt"
	"os"
	"strconv"

	"github.com/golang/glog"
	"github.com/spf13/viper"
)

func init() {
	SetSettingsFromViper()
}

//Settings : Contains all of the envs that need to be grabbed for this application

var (
	// ServerPort : port to run gin server on
	ServerPort int
	//ServerHostName : Hostname to run this server on
	ServerHostName string
	// Debug mode for otp messages
	Debug bool
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

	// Environment : dev environment, production, docker, etc
	Environment AppEnvironment

	// AppEnvironments : array of all app environments
	AppEnvironments = []AppEnvironment{
		AppEnvironmentTesting,
		AppEnvironmentLocal,
		AppEnvironmentStaging,
		AppEnvironmentProduction,
	}
)

// AppEnvironment : string wrapper for environment name
type AppEnvironment string

const (
	// AppEnvironmentTesting : testing env
	AppEnvironmentTesting = AppEnvironment("testing")
	// AppEnvironmentLocal :
	AppEnvironmentLocal = AppEnvironment("local")
	// AppEnvironmentStaging :
	AppEnvironmentStaging = AppEnvironment("staging")
	// AppEnvironmentProduction :
	AppEnvironmentProduction = AppEnvironment("production")
)

//SetSettingsFromViper : Sets global settings using viper
func SetSettingsFromViper() {
	Environment := getEnvironment()
	glog.Info("We're in our the following environment: ", Environment)
	// SetENV if not in a production environment
	// Check for local
	if Environment != AppEnvironmentProduction && Environment != AppEnvironmentStaging {
		setEnvironmentVariablesFromConfig(Environment)
	}

	if Environment == AppEnvironmentTesting {
		DbName = os.Getenv("TEST_DB_NAME")
	} else {
		DbName = os.Getenv("DB_NAME")
	}

	DbDriver = os.Getenv("DB_DRIVER")
	DbHostname = os.Getenv("DB_HOSTNAME")
	DbUsername = os.Getenv("DB_USERNAME")
	DbPort = os.Getenv("DB_PORT")
	DbName = os.Getenv("DB_NAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	glog.Info("Db settings: ", DbDriver, " ", DbHostname, " ", DbName)

	Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	ServerHostName = os.Getenv("SERVER_HOSTNAME")
	ServerPort, _ = strconv.Atoi(os.Getenv("SERVER_PORT"))
	if Environment != "staging" {

		ServerHostName = viper.GetString("serverHostName")
		ServerPort = viper.GetInt("serverPort")
		Debug = viper.GetBool("debug")

	}
}

func setEnvironmentVariablesFromConfig(env AppEnvironment) {
	// get and set basePath of project
	baseProjectPath := fmt.Sprintf("%s/src/<%= myrepoUrl %>/<%= myappName %>", os.Getenv("GOPATH"))
	viper.AddConfigPath(baseProjectPath + "/config/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("localConfig")

	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		glog.Info("Failed reading local settings: ", err)
	}
	debug := viper.GetBool("debug")

	serverHostName := viper.GetString("serverHostName")
	serverPort := viper.GetString("serverPort")
	dbDriver := viper.GetString("dbDriver")
	dbHostname := viper.GetString("dbHostName")
	dbPassword := viper.GetString("dbPassword")
	dbRootPassword := viper.GetString("dbRootPassword")
	dbPort := viper.GetString("dbPort")
	dbUser := viper.GetString("dbUsername")
	dbName := viper.GetString("dbName")
	dbTestDBName := viper.GetString("testDbName")

	// Set the OS Environment variables
	os.Setenv("DB_DRIVER", dbDriver)
	os.Setenv("DB_HOSTNAME", dbHostname)
	os.Setenv("DB_USERNAME", dbUser)
	os.Setenv("DB_PORT", dbPort)
	os.Setenv("DB_NAME", dbName)
	os.Setenv("DB_PASSWORD", dbPassword)
	os.Setenv("DB_ROOTPASSWORD", dbRootPassword)
	os.Setenv("TEST_DB_NAME", dbTestDBName)
	os.Setenv("DEBUG", strconv.FormatBool(debug))
	os.Setenv("SERVER_HOSTNAME", serverHostName)
	os.Setenv("SERVER_PORT", serverPort)
	glog.Info("setEnvironmentVariablesFromConfig: Config finished reading in settings from file.")

}

func getEnvironment() AppEnvironment {
	hostEnvironment := os.Getenv("SEZZLE_ENVIRONMENT")
	for _, env := range AppEnvironments {
		if env == AppEnvironment(hostEnvironment) {
			Environment = env
			return env
		}
	}

	// set to local config if environment not found
	return AppEnvironmentLocal
}

//IsProduction is a check for Production environment
func (e AppEnvironment) IsProduction() bool {
	return e == AppEnvironmentStaging || e == AppEnvironmentProduction
}
