package conf

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
	"os"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
}

func init() {
	go_path := os.Getenv("GOPATH")
	data, err := ioutil.ReadFile(go_path + "/bin/conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
