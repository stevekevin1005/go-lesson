package channel

import (
	"fmt"
	"time"
)

func Sample4() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = sample4WorkerCreater(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
		// n := <-channels[i]
	}
	time.Sleep(time.Millisecond)
}

func sample4WorkerCreater(id int) chan<- int {
	c := make(chan int)
	go func() {
		fmt.Printf("Worker %d received %c\n", id, <-c)
	}()
	return c
}
