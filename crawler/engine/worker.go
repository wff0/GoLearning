package engine

import (
	"GoBasic/crawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	//log.Printf("fetcher url: %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetcher fetching url: %s err: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.Parser.Parse(body, r.Url), nil
}
