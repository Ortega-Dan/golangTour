package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

	var actualFibo func(int) int

	actualFibo = func(numba int) int {
		switch numba {
		case 0:
			return 0
		case 1:
			return 1
		default:
			return actualFibo(numba-1) + actualFibo(numba-2)
		}
	}

	currentIndex := 0

	return func() int {
		currentIndex += 1

		return actualFibo(currentIndex)
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 12; i++ {
		fmt.Print(" ", f())
	}
}
