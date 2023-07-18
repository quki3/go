// slg = standard library go
package main

import (
	"fmt"
	"math"
)

// abs - Abs returns the absolute value of x.
func abs(req) {
	// math - slg
	res := math.Abs(req)
	fmt.Printf("%.1f\n", res)

	y := math.Abs(2)
	fmt.Printf("%.1f\n", y)
	return res
}
