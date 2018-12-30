package model

import (
	// _ "github.com/mattn/go-sqlite3"
	"log"
	"server/utils"
	"time"
)

type User struct {
	Id             int64     `xorm:"id"`
	Name           string    `xorm:"name"`
	Age            int       `xorm:"age"`
	Sex            string    `xorm:"sex"`
	HashedPassword string    `xorm:"password varchar(64)"`
	Mobile         string    `xorm:"mobile unique varchar(32)"`
	Created        time.Time `xorm:"created"`
	Updated        time.Time `xorm:"updated"`
	Password       string    `xorm:"-"`
}

func init() {
	err := db.Sync2(new(User))
	if err != nil {
		log.Fatalf("Fail to sync user: %v\n", err)
	}
}

func (u *User) New() bool {
	u.HashedPassword = utils.HashAndSalt([]byte(u.Password))

	_, err := db.Insert(u)
	if err != nil {
		log.Fatalf("Fail to create user: %v\n", err)
		return false
	}

	return true
}

func (u *User) QuerryUserByMobile() User {
	var user User
	_, err := db.Where("Mobile = ?", u.Mobile).Desc("id").Get(&user)

	if err != nil {
		log.Fatalf("Fail to query user: %v\n", err)
	}

	return user
}

func (u *User) HasRegistered() bool {
	var user User
	has, err := db.Where("Mobile = ?", u.Mobile).Desc("id").Get(&user)
	if err != nil {
		log.Fatalf("Fail to query user: %v\n", err)
	}

	if has {
		hashedPasswd := utils.HashAndSalt([]byte(u.Password))
		if utils.ComparePasswords(hashedPasswd, []byte(u.Password)) {
			return true
		}
	}

	return false
}
