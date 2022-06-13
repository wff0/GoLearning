package main

import (
	basicConfig "GoBasic/crawler/config"
	"GoBasic/crawler_distributed/config"
	"GoBasic/crawler_distributed/rpcsupport"
	"GoBasic/crawler_distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go serveRpc(host)
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764",
		Parser: worker.SerializedParser{
			Name: basicConfig.ParseProfile,
			Args: "寂寞成影萌宝",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
