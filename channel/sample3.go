package channel

import (
	"fmt"
	"time"
)

func Sample3() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = sample3WorkerCreater(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func sample3WorkerCreater(id int) chan int {
	c := make(chan int)
	go func() {
		fmt.Printf("Worker %d received %c\n", id, <-c)
	}()
	return c
}
