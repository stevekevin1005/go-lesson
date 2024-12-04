package engine

import (
	"go-examples/crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)
	itemCount := 0
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {

			log.Printf("Got article #%d: %v", itemCount, item)
			itemCount++

		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching: %s", r.Url)
	contents, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(contents), nil
}
