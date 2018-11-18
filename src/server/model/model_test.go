package model

import (
	"fmt"
	"testing"
)

func TestDb(t *testing.T) {
	orm := db

	_, err := orm.Insert(&User{Name: "xlw"})
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
