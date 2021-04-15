package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	if t != nil {

		if t.Left != nil {
			Walk(t.Left, ch)
		}

		ch <- t.Value

		if t.Right != nil {
			Walk(t.Right, ch)
		}
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	if t1 != nil && t2 != nil {
		same := true

		ch1 := make(chan int)
		ch2 := make(chan int)

		go walker(t1, ch1)
		go walker(t2, ch2)

		for v1 := range ch1 {

			v2 := <-ch2

			if v1 == v2 {
				same = false
			}
		}

		return same
	}

	return false

}

// convenience function to close the channel after
// the tree ends its recursive operation
func walker(tr *tree.Tree, ch chan int) {
	Walk(tr, ch)
	close(ch)
}

func main() {

	ch := make(chan int)

	// testing walk method
	go walker(tree.New(1), ch)
	// for range testing
	for v := range ch {
		fmt.Println(v)
	}

}
