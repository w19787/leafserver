package model

import (
	"flag"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"server/conf"
)

var db *xorm.Engine

func init() {
	if flag.Lookup("test.v") == nil {
		fmt.Println("testing mode")
		db = newEngine(conf.TestDBConfig)
	} else {
		db = newEngine(conf.ProductDBConfig)
	}
}

func newEngine(conf conf.DBEngineConfig) *xorm.Engine {
	db, err := xorm.NewEngine(conf.Driver, conf.DataSource)

	if err != nil {
		log.Fatalf("Fail to connect database: %v\n", err)
	}

	return db
}
