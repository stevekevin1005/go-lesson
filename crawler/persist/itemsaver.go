package persist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-examples/crawler/engine"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			Save(item)
			itemCount++
		}
	}()
	return out
}

func Save(item engine.Item) (id string, err error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		return "", err
	}

	// 測試連線有沒有通 (optional)
	_, err = es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return "", err
	}

	var buffer bytes.Buffer

	jsonData, err := json.Marshal(item.Payload)
	buffer.Write(jsonData)
	// if err := json.NewEncoder(&buffer).Encode(item.Payload); err != nil {
	// 	log.Fatalf("Error encoding item: %s", err)
	// 	return "", err
	// }

	res, err := es.Index(
		item.Site,
		&buffer,
		es.Index.WithRefresh("true"),
	)
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
		return "", err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error indexing document: %s", res)
		return "", fmt.Errorf("Error indexing document: %s", res)
	}

	var resBody map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
		return "", err
	}

	id, ok := resBody["_id"].(string)
	if !ok {
		log.Fatalf("Error parsing the response body: %s", err)
		return "", err
	}
	return id, nil
}
