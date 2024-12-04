package main

import (
	"go-examples/crawler/engine"
	"go-examples/crawler/persist"
	"go-examples/crawler/ptt/parser"
	"go-examples/crawler/scheduler"
)

// target: 獲取PTT文章列表
func main() {
	// engine.SimpleEngine{}.Run(engine.Request{
	// 	Url:        "https://www.ptt.cc/bbs/index.html",
	// 	ParserFunc: parser.ParseBoardList,
	// })
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 1000,
		ItemChan:    persist.ItemSaver(),
	}

	// e := engine.ConcurrentEngine{
	// 	Scheduler:   &scheduler.SimpleScheduler{},
	// 	WorkerCount: 100000,
	// 	ItemChan:    persist.ItemSaver(),
	// }
	e.Run(engine.Request{
		Url:        "https://www.ptt.cc/bbs/index.html",
		ParserFunc: parser.ParseBoardList,
	})
}
