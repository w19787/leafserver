package utils_test

import (
	"server/utils"
	"testing"
)

func TestPasswordString(t *testing.T) {

	pwd := "hello"
	s1 := utils.HashAndSalt([]byte(pwd))
	match := utils.ComparePasswords(s1, []byte(pwd))

	if !match {
		t.Error("password is not match")
	}
	s2 := utils.HashAndSalt([]byte("hello "))

	match = utils.ComparePasswords(s2, []byte(pwd))

	if match {
		t.Error("password is match")
	}

}

func TestPasswordNumber(t *testing.T) {

	pwd := "112233"
	s1 := utils.HashAndSalt([]byte(pwd))
	match := utils.ComparePasswords(s1, []byte(pwd))

	if !match {
		t.Error("password is not match")
	}

	pwd = "223344"
	match = utils.ComparePasswords(s1, []byte(pwd))

	if match {
		t.Error("password is match")
	}

}
