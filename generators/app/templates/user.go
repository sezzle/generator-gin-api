package gorm

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
}

// Save : Saves user object
func (u *User) Create() error {
	err := DB.Create(u).Error
	if err != nil {
		glog.Info(err)
		return err
	}

	return err
}
