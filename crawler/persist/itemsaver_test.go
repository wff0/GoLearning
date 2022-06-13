package persist

import (
	"GoBasic/crawler/engine"
	"GoBasic/crawler/model"
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
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

	const index = "dating_test"
	client, err := NewEsClient()
	if err != nil {
		panic(err)
	}
	err = Save(client, index, expected)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	es, err := NewEsClient()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	resp, err := es.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		log.Println(err)
		panic(err)
	}

	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	// Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
