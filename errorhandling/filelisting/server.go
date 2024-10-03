package main

import (
	"go-examples/errorhandling/filelisting/listing"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling request: %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(listing.HandleFileListing))
	// http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
	// 	path := request.URL.Path[len("/list/"):]
	// 	file, err := os.Open(path)
	// 	if err != nil {
	// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
	// 		return
	// 		// panic(err)
	// 	}
	// 	defer file.Close()
	// 	all, err := ioutil.ReadAll(file)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	writer.Write(all)
	// })
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
