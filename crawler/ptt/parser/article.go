package parser

import (
	"go-examples/crawler/engine"
	"go-examples/crawler/model"
	"regexp"
)

var (
	authorRe  = regexp.MustCompile(`<span class="article-meta-tag">作者<\/span><span class="article-meta-value">([^<]+)<\/span>`)
	titleRe   = regexp.MustCompile(`<span class="article-meta-tag">標題<\/span><span class="article-meta-value">([^<]+)<\/span>`)
	timeRe    = regexp.MustCompile(`<span class="article-meta-tag">時間<\/span><span class="article-meta-value">([^<]+)<\/span>`)
	contentRe = regexp.MustCompile(`<span class="article-meta-value">.*<\/span><\/div>([\s\S]+)--`)
	urlRe     = regexp.MustCompile(`<span class="f2">※ 文章網址: <a href="([^"]+)"`)
)

func ParseArticle(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	article := &model.Article{}
	if authorMatch := authorRe.FindSubmatch(contents); authorMatch != nil {
		article.Author = string(authorMatch[1])
	}
	if titleMatch := titleRe.FindSubmatch(contents); titleMatch != nil {
		article.Title = string(titleMatch[1])
	}
	if timeMatch := timeRe.FindSubmatch(contents); timeMatch != nil {
		article.Time = string(timeMatch[1])
	}
	if contentsMatch := contentRe.FindSubmatch(contents); contentsMatch != nil {
		article.Content = string(contentsMatch[1])
	}
	if urlReMatch := urlRe.FindSubmatch(contents); urlReMatch != nil {
		article.Url = string(urlReMatch[1])
	}
	item := engine.Item{
		Site:    "ptt",
		Url:     article.Url,
		Payload: article,
	}
	result.Items = append(result.Items, item)
	return result
}
