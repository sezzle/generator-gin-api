package gorm

import (
	"github.com/golang/glog"
)

//Migrate : Initializes all models in db
func Migrate() error {
	glog.Info("Running object Migrations...")

	/*
			===========================================
			Keep these alphabetical for easy search
		  ===========================================
	*/

	glog.Info("Creating User Table")
	err := DB.AutoMigrate(&User{}).Error
	if err != nil {
		glog.Info(err)
		return err
	}

	return err
}
