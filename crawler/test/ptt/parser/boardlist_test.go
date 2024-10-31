package test

import (
	"go-examples/crawler/ptt/parser"
	"os"
	"testing"
)

func TestParseBoardList(t *testing.T) {
	contents, err := os.ReadFile("boardlist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := parser.ParseBoardList(contents)
	const resultSize = 128
	expectedUrls := []string{
		"https://www.ptt.cc/bbs/Gossiping/index.html",
		"https://www.ptt.cc/bbs/LoL/index.html",
		"https://www.ptt.cc/bbs/C_Chat/index.html",
	}
	expectedBoardClass := []string{
		"綜合", "遊戲", "閒談",
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but got %s", i, url, result.Requests[i].Url)
		}
	}
	for i, boardClass := range expectedBoardClass {
		if result.Items[i] != boardClass {
			t.Errorf("expected board class #%d: %s; but got %s", i, boardClass, result.Items[i])
		}
	}
	if len(result.Requests) != resultSize {
		t.Errorf("expected requests size: %d; but got %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("expected items size: %d; but got %d", resultSize, len(result.Items))
	}
}
