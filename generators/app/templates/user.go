package gorm

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
}

// Save : Saves user object
func (u *User) Creates() error {
	err := DB.Create(u).Error
	if err != nil {
		glog.Info(err)
		return err
	}

	return err
}
