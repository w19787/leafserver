package model

import (
	"testing"
)

func TestUserNew(t *testing.T) {
	u := User{Mobile: "1234567", Name: "Alex", Age: 20, Password: "2345"}
	ret := u.New()

	if !ret {
		t.Error("create user error")
	}
}

func TestUserQueryByMobile(t *testing.T) {
	u := User{Mobile: "1234567", Name: "Alex", Age: 20, Password: "2345"}
	ret := u.New()

	if !ret {
		t.Error("create user error")
	}

	u1 := u.QuerryUserByMobile()
	if u1.Name != "Alex" {
		t.Error("Query user error")
	}
}
