package channel

import (
	"fmt"
	"time"
)

func Sample2() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go sample2Worker(i, channels[i])
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func sample2Worker(id int, c chan int) {
	for {
		n := <-c
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}
