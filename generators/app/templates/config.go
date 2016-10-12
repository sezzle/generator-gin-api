package config

import (
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
	// Environment : dev environment, production, docker, etc
	Environment string
)

//SetSettingsFromViper : Sets global settings using viper
func SetSettingsFromViper() {
	Environment := os.Getenv("ENVIRONMENT")

	//Check for docker defined environment variable
	if Environment == "docker" {
		//Try docker settings
		viper.SetConfigName("dockerConfig")
		viper.AddConfigPath("/go/src/sezzle/instantach/config/")
		viper.SetConfigType("yaml")

		err := viper.ReadInConfig()
		if err != nil {
			glog.Fatal("Could not properly load docker settings: ", err)
		}
	} else if Environment == "staging" {
		glog.Info("Reading from staging settings")

		ServerHostName = os.Getenv("SERVER_HOSTNAME")

		isDebug, err := strconv.ParseBool(os.Getenv("DEBUG"))
		if err != nil {
			glog.Info(err)
			isDebug = true
		}
		Debug = isDebug

	} else {
		//Try local settings first
		viper.SetConfigName("localConfig")
		viper.AddConfigPath("./config/")
		viper.SetConfigType("yaml")

		err := viper.ReadInConfig()
		if err != nil {
			glog.Info("Failed reading local settings: ", err)
		}
	}

	if Environment != "staging" {

		ServerHostName = viper.GetString("serverHostName")
		ServerPort = viper.GetInt("serverPort")
		Debug = viper.GetBool("debug")

	}
}
