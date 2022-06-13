package zhenai

import (
	"GoBasic/crawler/config"
	"GoBasic/crawler/engine"
	"regexp"
)

const cityListRegex = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[\da-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	rg := regexp.MustCompile(cityListRegex)

	result := engine.ParseResult{}
	matches := rg.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	return result
}
