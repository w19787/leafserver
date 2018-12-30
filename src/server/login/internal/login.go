package internal

import (
	"github.com/name5566/leaf/gate"
	"server/model"
	"server/msg"
)

func handleLogin(args []interface{}) {
	m := args[0].(*msg.LoginMsg)
	a := args[1].(gate.Agent)

	u := model.User{Mobile: m.Mobile, Password: m.Password}
	ret := u.HasRegistered()
	var cmsg msg.ClientMsg

	if !ret {
		cmsg = msg.ClientMsg{MsgId: 1, StatusCode: 401, Msg: "Incorrect UserName or Password"}
	} else {
		cmsg = msg.ClientMsg{MsgId: 1, StatusCode: 200, Msg: "success"}
	}
	a.WriteMsg(&cmsg)
}
