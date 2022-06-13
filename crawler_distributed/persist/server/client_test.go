package main

import (
	"GoBasic/crawler/engine"
	"GoBasic/crawler/model"
	"GoBasic/crawler_distributed/config"
	"GoBasic/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Url:  "http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764",
		Type: "zhenai",
		Id:   "8256018539338750764",
		Payload: model.Profile{
			Age:        83,
			Height:     105,
			Weight:     137,
			Income:     "财务自由",
			Gender:     "女",
			Name:       "寂寞成影萌宝",
			Xinzuo:     "狮子座",
			Occupation: "金融",
			Marriage:   "离异",
			House:      "无房",
			Hokou:      "南京市",
			Education:  "初中",
			Car:        "无车",
		},
	}
	var result string
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %v", result, err)
	}
}
