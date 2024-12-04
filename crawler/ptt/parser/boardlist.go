package parser

import (
	"go-examples/crawler/engine"
	"go-examples/crawler/ptt"
	"regexp"
)

// input: []byte encoded in UTF-8
func ParseBoardList(contents []byte) engine.ParseResult {
	// 1. 解析HTML
	re := regexp.MustCompile(`<a class="board" href="([^"]+)">[\s\S]*?<div class="board-class">([^<]+)</div>`)
	// 2. 找到PTT的文章列表
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	// 3. 印出來
	for _, match := range matches {
		// item := engine.Item{
		// 	Site:    "ptt",
		// 	Url:     ptt.PTT_URL + string(match[1]),
		// 	Payload: string(match[2]),
		// }
		// result.Items = append(result.Items, item)
		result.Requests = append(result.Requests, engine.Request{
			Url:        ptt.PTT_URL + string(match[1]),
			ParserFunc: ParseArticleList,
		})
	}
	return result
}
