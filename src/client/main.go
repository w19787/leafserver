package main

import (
	// "bufio"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network/protobuf"
	"net"
	// "os"
	"server/msg"
)

func encodeGameOpMsg(op int32) []byte {
	var Processor = protobuf.NewProcessor()
	Processor.Register(&msg.GameOp{})

	data, err := Processor.Marshal(&msg.GameOp{Op: op, Param: 88, Extra: "99"})
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

func decodeGameOpMsg(msgData []byte, n int) int32 {
	recv := &msg.GameOp{}
	err := proto.Unmarshal(msgData[4:n], recv)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	return recv.GetOp()
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	for {
		fmt.Print("Enter Op: ")
		var op int32
		_, err := fmt.Scanf("%d", &op)

		if err != nil {
			log.Fatal("read std error:", err)
		}

		m := encodeGameOpMsg(op)

		// 发送消息
		conn.Write(m)

		buf := make([]byte, 32)
		// 接收消息
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal("read error:", err)
		} else {
			fmt.Println(decodeGameOpMsg(buf, n))
		}
	}
}
