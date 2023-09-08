package src

import (
	"fmt"
	"math"
)

func Sqrt(x float64, args ...bool) float64 {
	debug := false
	if len(args) > 0 {
		debug = args[0]
	}

	z := 1.0

	diff := 1.0

	// loop this as long as diff is above a threshold
	for diff > 0.000000000001 {
		previousZ := z
		z -= (z*z - x) / (2 * z)

		if debug {
			fmt.Println(x, z)
		}

		diff = math.Abs(z - previousZ)
	}

	return z
}
