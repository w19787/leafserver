package login

import (
	"client/msg"
	"fmt"
)

func connectServer() bool {

}

func Login() {
	var mobile_phone string
	var password string

	_, err := fmt.Scanf("Mobile Phone: %s", &op)

	if err != nil {
		log.Fatal("Mobile input Error: ", err)
	}

	_, err = fmt.Scanf("Password: %s", &password)

	if err != nil {
		log.Fatal("Password input Error: ", err)
	}

}
