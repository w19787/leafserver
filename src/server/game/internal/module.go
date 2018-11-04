package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
	"fmt"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	fmt.Println("internal module onInit")
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}
