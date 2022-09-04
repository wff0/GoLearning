package main

import (
	"context"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	myserver "myserver/proto"
)

func main() {
	svr := service.New()

	client := myserver.NewMyserverService("myserver", svr.Client())

	hello, err := client.SayHello(context.Background(), &myserver.SayRequest{Message: "wangfengfan"})
	if err != nil {
		logger.Error(err)
	}
	logger.Info(hello.Answer)
}
