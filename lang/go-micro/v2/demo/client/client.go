package main

import (
	proto "GoBasic/lang/go-micro/v2/demo/proto/cap"
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
)

func main() {
	microV2()
}

func microV2() {
	// 实例化
	svr := micro.NewService(
		micro.Name("cap.fengfan.client"),
	)
	// 初始化
	svr.Init()

	capService := proto.NewCapService("cap.fengfan.service", svr.Client())

	response, err := capService.SayHello(context.TODO(), &proto.SayRequest{Message: "fengfan"})
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(response.Answer)
}
