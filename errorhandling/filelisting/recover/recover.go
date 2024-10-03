package main

import "errors"

func tryRecover() {

	type NetError struct {
		error
	}

	defer func() {
		r := recover()
		if err, ok := r.(NetError); ok {
			println("Error occurred:", err.Error())
		} else {
			panic(r)
		}
	}()
	panic(errors.New("this is an error"))
	// b := 0
	// a := 5 / b
	// println(a)
}

func main() {
	tryRecover()
}
