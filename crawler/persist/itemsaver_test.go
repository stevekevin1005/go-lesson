package persist

import (
	"encoding/json"
	"go-examples/crawler/model"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

func TestSave(t *testing.T) {
	article := &model.Article{
		Title: "[整理] 2024棒球賽事轉播節目表",
		Url:   "https://www.ptt.cc/bbs/Baseball/M.1704038825.A.DE5.html",
	}

	id, err := Save(article)
	if err != nil {
		panic(err)
	}

	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}

	res, err := client.Get("ptt_content", id)
	if err != nil {
		panic(err)
	}
	// 解析res回應，檢查是否有成功存入資料
	var doc map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		panic(err)
	}
	sourceMap, ok := doc["_source"].(map[string]interface{})
	if !ok {
		t.Error("expected _source to be map[string]interface{}, but got", doc["_source"])
	}
	//  目標: response -> json tring -> 自定義結構
	// 透過json.Marshal轉換成json string
	sourceJson, err := json.Marshal(sourceMap)
	if err != nil {
		panic(err)
	}
	// 透過json.Unmarshal轉換回結構
	var sourceArticle model.Article
	if err := json.Unmarshal(sourceJson, &sourceArticle); err != nil {
		panic(err)
	}

	if sourceArticle != *article {
		t.Error("expected article: ", *article, "but got", sourceArticle)
	}
}
