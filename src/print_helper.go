package src

import "fmt"

func PrettyPrint(f func(float64, ...bool) float64, x float64, debug bool) {
	fmt.Println("-------------------")

	s := f(x, debug)

	fmt.Printf("SQRT of %v is: %v \n", x, s)
	fmt.Println("-------------------")
}
