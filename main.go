package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"crawler/persist"
)

func main() {

	itemChan,err := persist.ItemSaver("dating_profile")
	if err!= nil {
		panic(err)
	}
	e := engine.ConncurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemChan:itemChan,
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,

	})
	//e.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc:parser.ParseCity,
	//})

}


