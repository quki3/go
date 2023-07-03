// slg = standard library go
package main

import (
	"fmt"
	"math"
)

// abs - Abs returns the absolute value of x.
func abs() {
	// math - slg
	x := math.Abs(-2)
	fmt.Printf("%.1f\n", x)

	y := math.Abs(2)
	fmt.Printf("%.1f\n", y)
}
