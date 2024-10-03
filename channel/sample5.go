package channel

func Sample5() {
	c := make(chan int)
	// go sample5Worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	// time.Sleep(time.Millisecond)
}

// func sample5Worker(id int, c chan int) {
// 	for {
// 		n := <-c
// 		fmt.Printf("Worker %d received %d\n", id, n)
// 	}
// }
