package msg

import (
	"encoding/binary"
	// "github.com/golang/protobuf/proto"
	"fmt"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network/protobuf"
	"reflect"
	"server/msg"
)

var Processor = protobuf.NewProcessor()

func msg_func(id uint16, t reflect.Type) {
	fmt.Println(id, t)
}

func init() {
	fmt.Println("client msg register")
	Processor.Register(&msg.ClientMsg{})
	Processor.Register(&msg.RegisterMsg{})
	Processor.Register(&msg.LoginMsg{})
	Processor.Register(&msg.GameOp{})
	fmt.Println("client msg register completed")

	Processor.Range(msg_func)
}

func EncodeMsg(msgBody interface{}) []byte {
	fmt.Println("start encoding msg: %T", msgBody)
	data, err := Processor.Marshal(msgBody)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println("encode success")

	// len + id + data
	m := make([]byte, 4+len(data[1]))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(2+len(data[1])))

	copy(m[2:], data[0])
	copy(m[4:], data[1])

	return m
}

func DecodeMsg(data []byte) (interface{}, error) {
	return Processor.Unmarshal(data)
}
