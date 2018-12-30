package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"server/model"
	"server/msg"
)

func handleRegister(args []interface{}) {
	m := args[0].(*msg.RegisterMsg)
	a := args[1].(gate.Agent)

	fmt.Println("receive client register req: %s", m.Name)

	u := model.User{Mobile: m.Mobile, Name: m.Name, Password: m.Password}
	ret := u.New()

	var cmsg msg.ClientMsg

	if ret {
		cmsg = msg.ClientMsg{MsgId: 0, StatusCode: 200, Msg: "success"}
		fmt.Println("Register success")
	} else {
		cmsg = msg.ClientMsg{MsgId: 0, StatusCode: 406, Msg: "Regiseter failed"}
		fmt.Println("Register failed")
	}

	a.WriteMsg(&cmsg)
}
