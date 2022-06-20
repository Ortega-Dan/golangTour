package B_concurrency

import (
	"fmt"
	"time"
)

func ChannelsBasic_main() {
	tick := time.Tick(1 * time.Second)
	boom := time.After(5 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
