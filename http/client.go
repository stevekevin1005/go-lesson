package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Accept-Encoding", "gzip")
	client := new(http.Client)
	response, err := client.Do(request)
	// response, err := http.DefaultClient.Do(request)
	// response, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	s, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
