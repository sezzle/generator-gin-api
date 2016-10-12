package gorm

import (
  	"github.com/jinzhu/gorm"
    "github.com/golang/glog"
)

type User struct {
  gorm.Model
  FirstName string
  LastName string
}

// Save : Saves user object
func (u *User) Save() error {
  err := DB.Create(u).Error
  if err != nil {
    glog.Info(err)
    return err
  }

  return err
}
