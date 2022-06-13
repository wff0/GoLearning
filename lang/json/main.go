package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID         string      `json:"id"`
	Items      []OrderItem `json:"item"`
	TotalPrice float64     `json:"total_price"`
}

func main() {
	parseNLP()
}

func marshal() {
	o := Order{
		ID:         "1",
		TotalPrice: 20,
		Items: []OrderItem{
			{
				ID:    "item_1",
				Name:  "learn go",
				Price: 15,
			},
			{
				ID:    "item_2",
				Name:  "interview",
				Price: 10,
			},
		},
	}

	//fmt.Printf("%+v\n", o)
	s, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}

func unmarshal() {
	s := `{"id":"1","item":[{"id":"item_1","name":"learn go","price":15},{"id":"item_2","name":"interview","price":10}],"total_price":20}`
	var o Order
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
}

func parseNLP() {
	res := `{
  "RequestId": "FA53D08F-37D1-4D81-BEE7-41F24E825F60",
  "Data": {
    "result": [
      {
        "synonym": "",
        "weight": "0.100000",
        "tag": "普通词",
        "word": "请"
      },
      {
        "synonym": "",
        "weight": "0.100000",
        "tag": "普通词",
        "word": "输入"
      },
      {
        "synonym": "",
        "weight": "1.000000",
        "tag": "品类",
        "word": "文本"
      }
    ],
    "success": true
  }
}`

	m := struct {
		RequestId string `json:"RequestId"`
		Data      struct {
			Success bool `json:"success"`
			Result  []struct {
				Synonym string `json:"synonym"`
				Weight  string `json:"weight"`
				Tag     string `json:"tag"`
				Word    string `json:"word"`
			} `json:"result"`
		} `json:"Data"`
	}{}
	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", m)

	//fmt.Printf("%+v\n", m["Data"].(string))
	//
	//m2 := map[string]interface{}{}
	//json.Unmarshal([]byte(m["Data"].(string)), &m2)
	//
	//fmt.Printf("%+v\n", m2["result"].([]interface{})[0].(map[string]interface{})["tag"])
}
