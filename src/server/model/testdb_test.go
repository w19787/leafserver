package dbConn

import (
	"fmt"
	// "os"
	"testing"

	"github.com/go-xorm/xorm"
	// _ "github.com/mattn/go-sqlite3"
	"server/conf"
)

// User describes a user
// type User struct {
// 	Id      int64
// 	Name    string
// 	Created time.Time `xorm:"created"`
// 	Updated time.Time `xorm:"updated"`
// }

func TestDb(t *testing.T) {
	// f := "file:test.db?cache=shared&mode=memory"
	// os.Remove(f)

	orm, err := xorm.NewEngine(conf.TestDBConfig.Driver, conf.TestDBConfig.DataSource)
	if err != nil {
		fmt.Println(err)
		return
	}
	// orm.ShowSQL(true)

	err = orm.CreateTables(&User{})
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = orm.Insert(&User{Name: "xlw"})
	if err != nil {
		fmt.Println(err)
		return
	}

	u := new(User)
	has, err := orm.ID(1).Get(u)
	fmt.Println(has, err)
	fmt.Println(u.Created)
	fmt.Println(u.Name)

	u2 := new(User)
	has, err = orm.ID(3).Get(u2)
	fmt.Println(has, err)
	fmt.Println(u2.Created)
	fmt.Println(u2.Name)
}
