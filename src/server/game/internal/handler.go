package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/msg"
	"fmt"
)

func init(){
	fmt.Println("internal handler init")
	handler(&msg.Hello{}, handleHello)
}

func handler(m interface{}, h interface {}){
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}){
	log.Debug("receive msg from client")
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)

	log.Debug("hello %v", m.GetName())

	a.WriteMsg(&msg.Hello{Name: proto.String("client back: hello")})
}
