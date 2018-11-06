package gate

import (
	"server/game"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.GameOp{}, game.ChanRPC)
}
