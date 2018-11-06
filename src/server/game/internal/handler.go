package internal

import (
	"fmt"
	// "github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/msg"
)

func init() {
	fmt.Println("internal handler init")
	handler(&msg.GameOp{}, handleGameOp)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleGameOp(args []interface{}) {
	log.Debug("receive msg from client")
	m := args[0].(*msg.GameOp)
	a := args[1].(gate.Agent)

	log.Debug("Op %v", m.GetOp())

	a.WriteMsg(&msg.GameOp{Op: m.GetOp()})
}
