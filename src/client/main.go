package main

import (
	"client/login"
	"fmt"
)

func main() {
	ret := login.Register()

	if ret {
		fmt.Println("Register success")
	} else {
		fmt.Println("Regsiter failed")
	}
}
