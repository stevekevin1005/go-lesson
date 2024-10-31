package test

import (
	"go-examples/crawler/model"
	"go-examples/crawler/ptt/parser"
	"os"
	"testing"
)

func TestParseArticle(t *testing.T) {
	contents, err := os.ReadFile("article_test_data.html")
	if err != nil {
		panic(err)
	}
	result := parser.ParseArticle(contents)

	expectArtile := &model.Article{
		Title: "[整理] 2024棒球賽事轉播節目表",
		Url:   "https://www.ptt.cc/bbs/Baseball/M.1704038825.A.DE5.html",
	}

	if resultItem, ok := result.Items[0].(*model.Article); ok {
		if resultItem.Title != expectArtile.Title {
			t.Error("expected title: ", expectArtile.Title, "but got", resultItem.Title)
		}
		if resultItem.Url != expectArtile.Url {
			t.Error("expected url: ", expectArtile.Url, "but got", resultItem.Url)
		}
	} else {
		t.Error("expected item to be Article type, but got", result.Items[0])
	}
}
