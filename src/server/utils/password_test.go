package utils_test

import (
	"server/utils"
	"testing"
)

func TestPassword(t *testing.T) {

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
