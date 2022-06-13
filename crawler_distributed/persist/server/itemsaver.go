package main

import (
	basicPersist "GoBasic/crawler/persist"
	"GoBasic/crawler_distributed/config"
	"GoBasic/crawler_distributed/persist"
	"GoBasic/crawler_distributed/rpcsupport"
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
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
	//err := serveRpc(":1234", "dating_profile")
	//if err != nil {
	//	panic(err)
	//}
}

func serveRpc(host, index string) error {
	client, err := basicPersist.NewEsClient()
	if err != nil {
		return err
	}
	err = rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
	if err != nil {
		return err
	}
	return nil
}
