package channel

import (
	"fmt"
	"time"
)

func Sample6() {
	c := make(chan int, 3)
	go sample6Worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
	time.Sleep(time.Millisecond)
}

func sample6Worker(id int, c chan int) {
	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}
