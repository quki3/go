package main
import (
	"fmt"
	"image/color"
)

func main(){
	var palette = []color.Color{color.White, color.Black}
	fmt.Println("First color:", palette[0])
	fmt.Println("Second color:", palette[1])
}
