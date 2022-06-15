package main

import (
	proto "GoBasic/lang/go-micro/proto/cap"
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
)

type CapService struct {
}

func (c *CapService) SayHello(ctx context.Context, request *proto.SayRequest, response *proto.SayResponse) error {
	response.Answer = "hello " + request.Message
	return nil
}

func main() {
	// 创建服务
	svr := micro.NewService(
		micro.Name("cap.fengfan.service"),
	)
	// 初始化方法
	svr.Init()
	// 注册服务
	if err := proto.RegisterCapHandler(svr.Server(), new(CapService)); err != nil {
		logger.Fatal(err)
	}
	// 运行服务
	if err := svr.Run(); err != nil {
		logger.Fatal(err)
	}
}
