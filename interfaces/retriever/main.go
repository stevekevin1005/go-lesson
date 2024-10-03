package main

import (
	"fmt"
	"go-examples/interfaces/retriever/mock"
	"go-examples/interfaces/retriever/real"
)

type RetreiverInterface interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post("http://www.google.com", map[string]string{
		"name":   "steve",
		"course": "golang",
	})
}

type RetreiverPoster interface {
	RetreiverInterface
	Poster
}

func session(s RetreiverPoster) string {
	s.Post("http://www.google.com", map[string]string{
		"name":   "steve",
		"course": "golang",
	})
	return s.Get("http://www.google.com")
}

func download(r RetreiverInterface) string {
	return r.Get("http://www.google.com")
}

func main() {
	// var r RetreiverInterface
	// r = mock.Retreiver{"this is a fake google.com"}
	// inspect(r)
	// r = real.Retreiver{
	// 	UserAgent: "Mozilla/5.0",
	// 	TimeOut:   time.Minute,
	// }
	// inspect(r)

	// fmt.Println(download(r))
	// fmt.Println("Try a session")
	fmt.Println(session(&mock.Retreiver{"this is a fake google.com"}))
	// Type assertion
	// realRetriever := r.(mock.Retreiver)
	// fmt.Println(realRetriever)
}

func inspect(r RetreiverInterface) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retreiver:
		fmt.Println("Contents:", v.Contents)
	case real.Retreiver:
		fmt.Println("UserAgent:", v.UserAgent)

	}
}
