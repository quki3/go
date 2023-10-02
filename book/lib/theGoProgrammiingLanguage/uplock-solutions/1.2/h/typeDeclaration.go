// Type declarations most often appear at package level, where the named type is visible throughout the package, and if the name is exported (it starts with an uppercase letter), it’s accessible from other packages as well. example: type Celsius float64
package main
import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC 	Celsius = -273.15
	FreezingC 	Celsius = 0
	BoilingC	Celsius = 100
	BoilingF	Fahrenheit = 212
)

func main (){
	var c Celsius
	var f Fahrenheit
	fmt.Println("celsius = 0? ",c == 0)
	fmt.Println("fahrenheit >= 0?",f >= 0)
	fmt.Println("celsius = fahrenheit?",c == f)
	fmt.Println("Celsius(f)",c == Celsius(f))


	fmt.Printf("boilingC - freezingC = %g\n", BoilingC - FreezingC)
	fmt.Printf("boilingF - FreezingC = %g\n", BoilingF - FreezingC) // Error type mismatch.
	fmt.Printf("boilingF - FreezingC = %g\n", boilingF - CToF(FreezingC))

}

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32)}
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}

