package conf

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type DBEngineConfig struct {
	Driver     string
	DataSource string
}

var ProductDBConfig = DBEngineConfig{
	Driver:     "mysql",
	DataSource: "root@/game_login?charset=utf8",
}

var TestDBConfig = DBEngineConfig{
	Driver:     "sqlite3",
	DataSource: "file:test.db?cache=shared&mode=memory",
}
