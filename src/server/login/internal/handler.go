package internal

import (
	"reflect"
	"server/msg"
)

func init() {
	handleMsg(&msg.RegisterMsg{}, handleRegister)
	handleMsg(&msg.LoginMsg{}, handleLogin)
}

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
