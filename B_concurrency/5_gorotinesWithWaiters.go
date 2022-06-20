package B_concurrency

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	defer stndWaiter.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(s)
	}
}

var stndWaiter sync.WaitGroup

func GoroutinesWithWaiters_main() {
	stndWaiter.Add(2)
	go say("world")
	say("hello")

	stndWaiter.Wait()
}
