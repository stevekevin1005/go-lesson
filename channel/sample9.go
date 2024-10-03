package channel

import (
	"fmt"
	"sync"
)

func Sample9() {
	var workers [10]worker9
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i] = createSample9Worker(i, &wg)
	}
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()
}

type worker9 struct {
	in   chan int
	done func()
}

func createSample9Worker(id int, wg *sync.WaitGroup) worker9 {
	w := worker9{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}

	go doSample9Work2(id, w)
	return w
}

func doSample9Work2(id int, w worker9) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}
