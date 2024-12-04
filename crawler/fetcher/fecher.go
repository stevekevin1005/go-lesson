package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

var rateLimiter = time.Tick(200 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	// 1. 知道網頁的編碼方式
	// encoding := determineEncoding(resp.Body)
	bodyReader := bufio.NewReader(resp.Body)
	encoding := determineEncoding(bodyReader)
	// 2. 用該編碼方式的解碼器解碼
	reader := transform.NewReader(bodyReader, encoding.NewDecoder())
	return io.ReadAll(reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		panic(err)
	}
	// 透過讀取網頁的部分內容，判斷網頁的編碼方式
	encoding, _, _ := charset.DetermineEncoding(bytes, "")
	return encoding
}
