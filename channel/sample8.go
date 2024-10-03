package channel

import (
	"fmt"
)

func Sample8() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createSample8Worker(i)
	}
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// wait for all workers to finish
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createSample8Worker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doSample8Work(id, w.in, w.done)
	return w
}

func doSample8Work(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		// done <- true
		go func() {
			done <- true
		}()
	}
}
