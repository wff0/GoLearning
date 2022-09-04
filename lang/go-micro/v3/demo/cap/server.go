package main

import (
	"GoBasic/lang/go-micro/v3/demo/proto/cap"
	"context"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

type CapService struct {
}

func (c *CapService) SayHello(ctx context.Context, req *proto.SayRequest, res *proto.SayResponse) error {
	res.Answer = "你也好 " + req.Message
	return nil
}

func main() {
	microV3()
}

func microV3() {
	svr := service.New(service.Name("cap"))

	svr.Init()

	if err := proto.RegisterCapHandler(svr.Server(), new(CapService)); err != nil {
		logger.Fatal(err)
	}

	if err := svr.Run(); err != nil {
		logger.Fatal(err)
	}
	logger.Info("Register success")
}

func microV2() {
	//// 创建服务
	//svr := micro.NewService(
	//	micro.Name("cap.fengfan.service"),
	//)
	//// 初始化方法
	//svr.Init()
	//// 注册服务
	//if err := proto.RegisterCapHandler(svr.Server(), new(CapService)); err != nil {
	//	logger.Fatal(err)
	//}
	//// 运行服务
	//if err := svr.Run(); err != nil {
	//	logger.Fatal(err)
	//}
}
