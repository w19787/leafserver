package model

import (
	"log"
	"testing"
)

// func TestMain(m *testing.M) {

// 	exitVal := m.Run()
// 	log.Println("drop tables")
// 	db.DropTables()

// 	os.Exit(exitVal)
// }

func setupTest(t *testing.T) func() {
	// Test setup
	log.Println("setupTest()")
	db.CreateTables(new(User))

	// Test teardown - return a closure for use by 'defer'
	return func() {
		// t is from the outer setupTest scope
		db.DropTables("User")
		log.Println("teardownTest()")
	}
}

func TestUserNew(t *testing.T) {
	defer setupTest(t)()
	log.Println("excuting")
	u := User{Mobile: "1234567", Name: "Alex", Age: 20, password: "2345"}
	ret := u.New()

	if !ret {
		t.Error("create user error")
	}
}

func TestUserQueryByMobile(t *testing.T) {
	defer setupTest(t)()
	u := User{Mobile: "1234567", Name: "Alex", Age: 20, password: "2345"}
	ret := u.New()

	if !ret {
		t.Error("create user error")
	}

	u1 := u.QuerryUserByMobile()
	if u1.Name != "Alex" {
		t.Error("Query user error")
	}
}

func TestHasRegister(t *testing.T) {
	defer setupTest(t)()
	u := User{Mobile: "1234567", Name: "Alex", Age: 20, password: "2345"}
	u.New()

	ret := u.HasRegistered()
	if !ret {
		t.Error("user has not been registered")
	}
}
