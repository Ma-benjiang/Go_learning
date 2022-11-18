package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	count := 1
	// for i:=0;i<10;i++{
	// 	fmt.Println(z,z*z-x)
	// 	z -= (z*z - x) / (2 * z)
	// }
	for math.Abs(z*z-x) > 4.5e-16 {
		// fmt.Println(z)
		z -= (z*z - x) / (2 * z)
		count++
	}
	fmt.Println(count)
	return z
}
func main() {
	fmt.Println(Sqrt(3))
}
