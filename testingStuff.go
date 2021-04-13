package main

import "fmt"

// import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	y := make([][]uint8, dy)

	for yi := range y {
		y[yi] = make([]uint8, dx)

		for xi := range y[yi] {
			// y[yi][xi] = uint8((xi + yi) / 2)
			// y[yi][xi] = uint8(xi * yi)
			y[yi][xi] = uint8(xi ^ yi)
		}
	}
	return y
}

func main() {
	// pic.Show(Pic)
	for _, v := range Pic(4, 5) {
		fmt.Println(v)
	}
}
