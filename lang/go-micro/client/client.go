package main

import (
	proto "GoBasic/lang/go-micro/proto/cap"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
)

func main() {
	// 实例化
	svr := micro.NewService(
		micro.Name("cap.fengfan.service"),
	)
	// 初始化
	svr.Init()

	capService := proto.NewCapService("cap.fengfan.service", svr.Client())

	response, err := capService.SayHello(context.Background(), &proto.SayRequest{Message: "fengfan"})
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(response.Answer)
}
