package main

import (
	"bufio"
	"fmt"
	"os"
)

// func tryDefer() {
// 	defer println(1)
// 	defer println(2)
// 	println(3)
// }

// func tryDefe2() {
// 	for i := 0; i < 100; i++ {
// 		defer println(i)
// 		if i == 30 {
// 			panic("printed too many")
// 		}
// 	}
// }

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func writeFile(filename string) {
	// file, err := os.Create(filename)

	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	// err = errors.New("this is a customer error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			fmt.Println("Unknown error:", err)
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// tryDefer()
	// // tryDefe2()
	writeFile("fib.txt")
}
