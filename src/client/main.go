package main

import (
	"client/login"
	"fmt"
)

func main() {
	fmt.Println("Chose function...")
	fmt.Println("	1: Register")
	fmt.Println("	2: Login")

	var inputFunc int
	var ret bool

	fmt.Scanln(&inputFunc)

	if inputFunc == 1 {
		ret = login.Register()
	} else {
		ret = login.Login()
	}

	if ret {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
}
