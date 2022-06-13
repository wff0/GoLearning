package zhenai

import (
	"GoBasic/crawler/config"
	"GoBasic/crawler/engine"
	"regexp"
)

// <a href="http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764">寂寞成影萌宝</a>
//const cityRegex = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>`

var (
	profileRe = regexp.MustCompile(`<a href="(.*album\.zhenai\.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(.*www\.zhenai\.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	result := engine.ParseResult{}
	matches := profileRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//result.Items = append(result.Items, fmt.Sprintf("User %s", m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	return result
}
