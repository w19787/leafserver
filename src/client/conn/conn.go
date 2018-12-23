package conn

import (
	"net"
	"strconv"
	"fmt"
)

type LoginServer struct{
	ServerType string
	Server string
	Port int
	conn net.Conn
}

func (s *LoginServer) Connect() bool{
	host := s.Server + strconv.Itoa(s.Port)
	s.conn, err := net.Dial(s.ServerType, host)

	if err != nil{
		fmt.Println(err)
		return false
	}

	return true
}

func (s *LoginServer) Write(data []byte){
	s.conn.Write(data)
}

func (s *LoginServer) Read(buf []byte) (int, error){
	return s.conn.Read(buf)
}

