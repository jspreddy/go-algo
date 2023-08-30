package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	diff := 1.0

	// loop this as long as diff is above a threshold
	for diff > 0.000000000001 {
		previousZ := z
		z -= (z*z - x) / (2 * z)
		fmt.Println(x, z)
		diff = math.Abs(z - previousZ)
	}
	return z
}

func prettyPrint(f func(float64) float64, x float64) {
	fmt.Println("-------------------")

	s := f(x)

	fmt.Printf("SQRT of %v is: %v \n", x, s)
	fmt.Println("-------------------")
}

func main() {
	// prettyPrint(Sqrt, 2)
	// prettyPrint(Sqrt, 4)
	prettyPrint(Sqrt, 144)
}
