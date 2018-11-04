package main

import (
	"server/msg"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/network/protobuf"
	"github.com/name5566/leaf/log"
	"net"
	"bufio"
	"os"
)

func encodeHelloMsg(msgBody string) []byte{
	var Processor = protobuf.NewProcessor()
	Processor.Register(&msg.Hello{})

	data, err := Processor.Marshal(&msg.Hello{Name: proto.String(msgBody)})
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

func decodeHelloMsg(msgData []byte, n int) string{
	recv := &msg.Hello{}
	err := proto.Unmarshal(msgData[4:n], recv)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	return recv.GetName()
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	for{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')

		m := encodeHelloMsg(text)

		// 发送消息
		conn.Write(m)

		buf := make([]byte, 32)
		// 接收消息
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal("read error:", err)
		}else{
			fmt.Println(decodeHelloMsg(buf, n))
		}
	}
}
