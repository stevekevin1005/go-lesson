package goroutine

import (
	"fmt"
	"time"
)

func Sample3() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
	fmt.Println("Done")
}
