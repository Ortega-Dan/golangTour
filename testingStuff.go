package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Ortega-Dan/golang-stdin/stdin"
)

func Sqrt(x float64) float64 {

	lastZ := 0.0

	z := 200.0

	for i := 0; true; i++ {
		z -= (z*z - x) / (2 * z)

		fmt.Printf("Iteration %v, with result %v\n", i, z)
		if math.Abs(z-lastZ) < 0.000001 {
			fmt.Println()
			return z
		}
		lastZ = z
	}
	fmt.Println()
	return z
}

func main() {

	fmt.Print("Enter number to calculate the Square Root: ")

	numba, _ := strconv.ParseFloat(stdin.ReadLine(), 64)

	fmt.Println()
	fmt.Printf("%1.4f\n", Sqrt(numba))

	stdin.ReadLine()

}
