package main

import (
	"GoBasic/crawler_distributed/rpcsupport"
	"GoBasic/crawler_distributed/worker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port)))
}

func serveRpc(host string) error {
	err := rpcsupport.ServeRpc(host, worker.CrawlService{})
	if err != nil {
		return err
	}
	return nil
}
