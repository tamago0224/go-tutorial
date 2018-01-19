package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func Sqrt2(x float64) float64 {
	z_delta, z := x, 1.0
	for math.Abs(z-z_delta) >= 1.0e-6 {
		z_delta, z = z, z-(z*z-x)/(2*z)
	}
	return z
}
func main() {
	fmt.Println(Sqrt2(2))
}
