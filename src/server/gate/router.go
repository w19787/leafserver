package gate

import (
	"server/game"
	"server/login"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.RegisterMsg{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.LoginMsg{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.GameOp{}, game.ChanRPC)
}
