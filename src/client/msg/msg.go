package msg

import (
	"encoding/binary"
	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(&ClientMsg{})
	Processor.Register(&RegisterMsg{})
	Processor.Register(&LoginMsg{})
	Processor.Register(&GameOp{})
}

func EncodeMsg(msgBody interface{}) []byte {
	data, err := Processor.Marshal(msgBody.(proto.Message))
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

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
