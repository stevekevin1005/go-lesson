package goroutine

import (
	"fmt"
	"time"
)

func Sample1() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Second)
}
