package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	defer waiter.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(s)
	}
}

var waiter sync.WaitGroup

func main() {
	waiter.Add(2)
	go say("world")
	say("hello")

	waiter.Wait()
}
