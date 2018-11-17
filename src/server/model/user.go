package dbConn

import (
	// _ "github.com/mattn/go-sqlite3"
	"time"
)

type User struct {
	Id       int64
	Name     string
	Age      int
	Password string    `xorm:"varchar(32)"`
	Mobile   string    `xorm:"varchar(32)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}
