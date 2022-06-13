package zhenai

import (
	"GoBasic/crawler/engine"
	"GoBasic/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := parseProfile(contents,
		"http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764",
		"寂寞成影萌宝")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	actual := result.Items[0]
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

	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}
}
