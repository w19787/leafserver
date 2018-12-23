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
	cmsg "client/msg"
	"server/msg"
)

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
