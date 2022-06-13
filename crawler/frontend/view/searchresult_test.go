package view

import (
	"GoBasic/crawler/engine"
	"GoBasic/crawler/frontend/model"
	common "GoBasic/crawler/model"
	"os"
	"testing"
)

func TestSearchResultView(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")
	page := model.SearchResult{}
	page.Hits = 1234
	expected := engine.Item{
		Url:  "http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764",
		Type: "zhenai",
		Id:   "8256018539338750764",
		Payload: common.Profile{
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
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, expected)
	}

	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
