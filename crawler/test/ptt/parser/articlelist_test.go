package test

import (
	"go-examples/crawler/model"
	"go-examples/crawler/ptt/parser"
	"os"
	"testing"
)

func TestParseArticleList(t *testing.T) {
	contents, err := os.ReadFile("articlelist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := parser.ParseArticleList(contents)
	expectedUrls := []string{
		"https://www.ptt.cc/bbs/Baseball/index17578.html",
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but got %s", i, url, result.Requests[i].Url)
		}
	}
	expectedArticles := []model.Article{
		{Title: "[討論] 非科班數據人才當中華隊總仔", Url: "https://www.ptt.cc/bbs/Baseball/M.1730297864.A.0EF.html"},
		{Title: "[分享] 田中將大今天在洋基球場看球", Url: "https://www.ptt.cc/bbs/Baseball/M.1730298605.A.58A.html"},
	}
	for i, article := range expectedArticles {
		if resultItem, ok := result.Items[i].(*model.Article); ok {
			if resultItem.Title != article.Title {
				t.Error("expected title: ", article.Title, "but got", resultItem.Title)
			}
			if resultItem.Url != article.Url {
				t.Error("expected url: ", article.Url, "but got", resultItem.Url)
			}
		} else {
			t.Error("expected item to be Article type, but got", result.Items[i])
		}
	}
}
