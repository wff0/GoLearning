package persist

import (
	"GoBasic/crawler/engine"
	"context"
	"errors"
	"log"

	"github.com/olivere/elastic/v7"
)

//var es *elastic.Client

//func init() {
//	var err error
//	es, err = NewEsClient()
//	if err != nil {
//		panic(err)
//	}
//}

const EsUrl = "http://192.168.221.128:9200"

func NewEsClient() (*elastic.Client, error) {
	return elastic.NewClient(
		elastic.SetURL(EsUrl),
		elastic.SetSniff(false))
}

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := NewEsClient()
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: got item #%d: %v", itemCount, item)
			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}

			itemCount++
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
