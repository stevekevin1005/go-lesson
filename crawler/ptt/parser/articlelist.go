package parser

import (
	"go-examples/crawler/engine"
	"go-examples/crawler/model"
	"go-examples/crawler/ptt"
	"regexp"
)

var (
	prevRe        = regexp.MustCompile(`<a class="btn wide" href="([^"]+)">.*上頁<\/a>`)
	articleListRe = regexp.MustCompile(`<div class="title">[\s\S]*?<a href="([^"]+)">(.+)<\/a>[\s\S]*?<\/div>`)
)

func NewArticle(title string, url string) *model.Article {
	return &model.Article{
		Title: title,
		Url:   url,
	}
}

// input: []byte encoded in UTF-8
func ParseArticleList(contents []byte) engine.ParseResult {
	// 1. 解析HTML 找到文章列表的上一頁 URL
	result := engine.ParseResult{}
	prevUrl := prevRe.FindSubmatch(contents)
	if prevUrl != nil {
		result.Requests = append(result.Requests, engine.Request{
			Url:        ptt.PTT_URL + string(prevUrl[1]),
			ParserFunc: ParseArticleList,
		})
	}
	// 2. 解析HTML 找到文章列表的文章 URL
	matches := articleListRe.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		url := ptt.PTT_URL + string(match[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: ParseArticle,
		})
		result.Items = append(result.Items, NewArticle(string(match[2]), url))
	}
	return result
}
