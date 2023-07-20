package main
import (
	"fmt"
	"os"
	"strconv"
	"/mnt/chroot/home/cloud/hub/go/uplock-solutions/1.2/i/gopl.io/ch2/conv"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC 	Celsius = -273.15
	FreezingC	Celsius = 0
	BoilingC	Celsius = 100
)
func main()	{
	func fc(c Celsius) String() string { return fmt.Springf("%g°C",c)}
	func ff(f Fahrenheit) String() string {return fmt.Sprintf("%g°F",f)}
//test
	var tempC Celsius = 70.7
	var tempF Fahrenheit = 140.7
	fmt.Println(fc(tempC))
	fmt.Println(ff(tempF))
	fmt.Println(conv.CToF(BoilingC))
}
