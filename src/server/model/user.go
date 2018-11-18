package model

import (
	// _ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type User struct {
	Id       int64
	Name     string
	Age      int
	Sex      string
	Password string    `xorm:"varchar(32)"`
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
