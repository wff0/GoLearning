package main

import (
	"GoBasic/lang/go-micro/v3/demo/proto/cap"
	"context"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	microV3()
}

func microV3() {
	svr := service.New()

	capService := proto.NewCapService("cap", svr.Client())

	res, err := capService.SayHello(context.Background(), &proto.SayRequest{Message: "wff"})
	if err != nil {
		logger.Debug(err)
	}
	logger.Info(res.Answer)
}

func microV2() {
	//// 实例化
	//svr := micro.NewService(
	//	micro.Name("cap.fengfan.service"),
	//)
	//// 初始化
	//svr.Init()
	//
	//capService := proto.NewCapService("cap.fengfan.service", svr.Client())
	//
	//response, err := capService.SayHello(context.Background(), &proto.SayRequest{Message: "fengfan"})
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//fmt.Println(response.Answer)
}
