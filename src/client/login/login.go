package login

import (
	"client/conn"
	cmsg "client/msg"
	"fmt"
	"github.com/name5566/leaf/log"
	smsg "server/msg"
)

func connectServer() *conn.Server {
	s := new(conn.Server)
	s.ServerType = "tcp"
	s.Server = "127.0.0.1"
	s.Port = 3563

	for i := 0; i < 3; i++ {
		if s.Connect() {
			fmt.Println("connect login server success")
			return s
		} else {
			fmt.Printf("try %d time", i+1)
			break
		}
	}

	return nil
}

func authenticate(s *conn.Server, mp string, password string) bool {

	loginMsg := smsg.LoginMsg{Mobile: mp, Password: password}
	s.Write(cmsg.EncodeMsg(&loginMsg))
	buf := make([]byte, 128)
	n, err := s.Read(buf)

	if err != nil {
		log.Fatal("read error:", err)
	}

	rsp, err := cmsg.DecodeMsg(buf[2:n])

	rsp_msg := rsp.(*smsg.ClientMsg)

	if err == nil && rsp_msg.StatusCode == 200 {
		return true
	}

	return false
}

func Login() bool {
	var mobile_phone string
	var password string

	fmt.Println("Mobile Phone: ")
	_, err := fmt.Scanf("%s", &mobile_phone)

	if err != nil {
		log.Fatal("Mobile input Error: ", err)
	}

	fmt.Println("Password: ")
	_, err = fmt.Scanf("%s", &password)

	if err != nil {
		log.Fatal("Password input Error: ", err)
	}

	s := connectServer()
	if s != nil {
		return authenticate(s, mobile_phone, password)
	}

	return false
}
