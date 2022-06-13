package main

import (
	"GoBasic/crawler/config"
	"GoBasic/crawler/engine"
	"GoBasic/crawler/persist"
	"GoBasic/crawler/scheduler"
	"GoBasic/crawler/zhenai/zhenai"
)

func main() {
	saver, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         saver,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:    "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(zhenai.ParseCityList, config.ParseCityList),
	})
	//e.Run(engine.Request{
	//	Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
