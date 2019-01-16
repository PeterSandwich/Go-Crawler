package engine

import (
	"log"
	"crawler/fetcher"
)

func worker(r Request) (ParseResult,error){
	log.Printf("Fetching %s",r.Url)
	body,err := fetcher.Fetch(r.Url)
	if err!=nil {
		log.Printf("Error in fetching url %s: %v\n",r.Url,err)
		return ParseResult{},err
	}
	return r.ParserFunc(body,r.Url),nil
}
