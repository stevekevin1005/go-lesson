package channel

import (
	"fmt"
	"time"
)

func Sample7() {
	c := make(chan int, 3)
	go sample7Worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
	time.Sleep(time.Millisecond)
}

func sample7Worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}
