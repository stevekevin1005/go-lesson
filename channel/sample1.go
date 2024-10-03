package channel

import "fmt"

func Sample1() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	// n := <-c
	// fmt.Println(n)
}
