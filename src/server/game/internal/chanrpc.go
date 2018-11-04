package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
)

func init() {
	fmt.Println("internal chanrpc init")
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	fmt.Println("internal chanrpc rpcNewAgent")
	a := args[0].(gate.Agent)
	fmt.Println(a.LocalAddr(), a.RemoteAddr())
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}
