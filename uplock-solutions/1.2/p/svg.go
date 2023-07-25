package main
import (
	"fmt"
	"math"
)

constn (
	width, height 	= 600, 320
	cells		= 100
	xyrange		= 30.0
	xyscale		= width / 2 / xyrange
	zscale		= height * 0.4
	angle		= math.Pi / 6
)
var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main(){
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		   "style='stroke: grey; fill: white; stroke-width: 0.7' "+
		   "width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1,j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1,j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g %g' />\n",
				   ax, ay, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg")
}
	
