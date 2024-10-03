package goroutine

import (
	"fmt"
	"time"
)

func Sample2() {
	var a [10]int
	for i := 0; i < 10; i++ {
		// go func() {
		go func(i int) {
			for {
				a[i]++
				// runtime.Gosched()
			}
		}(i)
		// }()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
