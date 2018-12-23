package msg

import (
	"fmt"
	"github.com/name5566/leaf/network/protobuf"
	"reflect"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(&ClientMsg{})
	Processor.Register(&RegisterMsg{})
	Processor.Register(&LoginMsg{})
	Processor.Register(&GameOp{})
}
