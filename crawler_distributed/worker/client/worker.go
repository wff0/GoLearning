package client

import (
	"GoBasic/crawler/engine"
	"GoBasic/crawler_distributed/config"
	"GoBasic/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(r engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(r)

		var sResult worker.ParseResult

		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(sResult), nil
	}
}
