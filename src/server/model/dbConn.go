package dbConn

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

func NewEngine(dbConf string) *xorm.Engine {
	db, err := xorm.NewEngine("mysql", dbConf)

	if err != nil {
		log.Fatalf("Fail to connect database: %v\n", err)
	}

	return db
}
