package db

import (
	"fmt"

        utils "github.com/mrtyunjaygr8/pass-n-go/utils"

	"gorm.io/gorm"
)

type User struct {
  gorm.Model
  Username string
  Password string
}

func (u *User) CreateUser(db *gorm.DB) (err error) {
  var count int64
  db.Model(&User{}).Where("Username = ?", u.Username).Count(&count)
  if count != 0 {
    return fmt.Errorf(utils.ALREADY_EXISTS)
  }
  result := db.Create(u)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

