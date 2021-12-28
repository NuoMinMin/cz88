package rpc

import (
	"cz88/pb"
	"fmt"
	"net"

	"cz88/internal/service"

	"cz88/config"

	"google.golang.org/grpc"
)

func New() {
	s := grpc.NewServer()
	pb.RegisterAppServer(s, service.New())
	listener, err := net.Listen("tcp", config.GetInstance().Rpc)
	if err != nil {
		panic("服务监听端口失败" + err.Error())
	}

	fmt.Println("rpc server addr: ", config.GetInstance().Rpc)
	if err = s.Serve(listener); err != nil {
		panic("服务启动失败" + err.Error())
	}
}
