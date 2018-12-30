package login

import (
	"client/conn"
	cmsg "client/msg"
	"fmt"
	"github.com/name5566/leaf/log"
	smsg "server/msg"
)

func online_register(s *conn.Server, name string, mp string, password string) bool {
	regMsg := smsg.RegisterMsg{Name: name, Mobile: mp, Password: password}

	fmt.Println("send Register info to server")
	s.Write(cmsg.EncodeMsg(&regMsg))
	fmt.Println("waiting for server rsp")

	buf := make([]byte, 128)
	n, err := s.Read(buf)

	if err != nil {
		log.Fatal("read error:", err)
	}
	fmt.Println("received server rsp, decoding...")
	rsp, err := cmsg.DecodeMsg(buf[2:n])

	client_msg := rsp.(smsg.ClientMsg)

	if err == nil && client_msg.StatusCode == 200 {
		return true
	}

	return false
}

func Register() bool {
	var name string
	var mobile_phone string
	var password string

	fmt.Print("User Name: ")
	_, err := fmt.Scanln(&name)

	if err != nil {
		log.Fatal("Name input Error: ", err)
	}

	fmt.Print("Mobile Phone: ")
	_, err = fmt.Scanln(&mobile_phone)

	if err != nil {
		log.Fatal("Mobile input Error: ", err)
	}

	fmt.Print("Password: ")
	_, err = fmt.Scanln(&password)

	if err != nil {
		log.Fatal("Password input Error: ", err)
	}

	s := connectServer()
	if s != nil {
		return online_register(s, name, mobile_phone, password)
	}

	return false
}
