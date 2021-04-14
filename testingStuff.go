package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	lastZ := 0.0

	z := 200.0

	for i := 0; true; i++ {
		z -= (z*z - x) / (2 * z)

		fmt.Printf("Iteration %v, with result %v\n", i, z)
		if math.Abs(z-lastZ) < 0.000001 {
			// fmt.Println()
			return z, nil
		}
		lastZ = z
	}
	// fmt.Println()
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
