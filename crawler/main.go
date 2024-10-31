package main

import (
	"go-examples/crawler/engine"
	"go-examples/crawler/ptt/parser"
)

// target: 獲取PTT文章列表
func main() {
	engine.Run(engine.Request{
		Url:        "https://www.ptt.cc/bbs/index.html",
		ParserFunc: parser.ParseBoardList,
	})
}
