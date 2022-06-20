package A_basicLangFeatures

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func picc(dx, dy int) [][]uint8 {

	y := make([][]uint8, dy)

	for yi, _ := range y {
		y[yi] = make([]uint8, dx)

		for xi := range y[yi] {
			// y[yi][xi] = uint8((xi + yi) / 2)
			// y[yi][xi] = uint8(xi * yi)
			y[yi][xi] = uint8(xi ^ yi)
		}
	}
	return y
}

func Ranges_main() {

	pic.Show(picc)

	println()

	for _, v := range picc(4, 5) {
		fmt.Println(v)
	}

	println("\nINFO: convert the base64 output image above to a png by doing:\necho 'theBase64string' | base64 -d > theFilenameToSave.png")
}
