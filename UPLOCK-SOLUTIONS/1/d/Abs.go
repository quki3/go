package main

import (
	"fmt"
	"math"
)

func abs() {
	x := math.Abs(-2)
	fmt.Printf("%.1f\n", x)

	y := math.Abs(2)
	fmt.Printf("%.1f\n", y)
}
