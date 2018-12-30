package conn

import (
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	ServerType string
	Server     string
	Port       int
	conn       net.Conn
}

func (s *Server) Connect() bool {
	host := s.Server + ":" + strconv.Itoa(s.Port)
	conn, err := net.Dial(s.ServerType, host)

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		s.conn = conn
	}

	return true
}

func (s *Server) Write(data []byte) {
	s.conn.Write(data)
}

func (s *Server) Read(buf []byte) (int, error) {
	return s.conn.Read(buf)
}
