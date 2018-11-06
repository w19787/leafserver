package base

import (
	"fmt"
	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/module"
	"server/conf"
)

func NewSkeleton() *module.Skeleton {
	fmt.Println("new Skeleton")
	skeleton := &module.Skeleton{
		GoLen:              conf.GoLen,
		TimerDispatcherLen: conf.TimerDispatcherLen,
		AsynCallLen:        conf.AsynCallLen,
		ChanRPCServer:      chanrpc.NewServer(conf.ChanRPCLen),
	}
	fmt.Println("NewSkeleton handler: ", skeleton)
	skeleton.Init()
	return skeleton
}
