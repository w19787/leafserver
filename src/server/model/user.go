package model

import (
	// _ "github.com/mattn/go-sqlite3"
	"log"
	"server/utils"
	"time"
)

type User struct {
	Id       int64
	Name     string
	Age      int
	Sex      string
	Password string    `xorm:"varchar(64)"`
	Mobile   string    `xorm:"varchar(32)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

func init() {
	err := db.Sync2(new(User))
	if err != nil {
		log.Fatalf("Fail to sync user: %v\n", err)
	}
}

func (u *User) New() bool {
	passwd := u.Password
	u.Password = utils.HashAndSalt([]byte(passwd))

	_, err := db.Insert(u)
	if err != nil {
		log.Fatalf("Fail to create user: %v\n", err)
		return false
	}

	return true
}

func (u *User) QuerryUserByMobile() User {
	user := User{}
	_, err := db.Where("Mobile = ?", u.Mobile).Desc("id").Get(&user)

	if err != nil {
		log.Fatalf("Fail to query user: %v\n", err)
	}

	return user
}
