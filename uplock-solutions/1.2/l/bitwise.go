package main
import "fmt"

//lexical variables
var z = uint(0)
var a uint8 = 1
var b uint8 = a<<1
var c uint8 = 5
var x uint8 = a<<1 | 1<<c
var y uint8 = a<<1 | 1<<2

func main() {
	fmt.Printf("%08b\n",z)
	fmt.Printf("type data=>[08b = binario, d = decimal] variables=>[a=1,b=a<<1,x=a<<a|a<<c] \n(a)08b = %08b (a)d = %d\n",a,a)	
	fmt.Printf("(b)08b = %08b (b)d = %d\n",b,b)			// <<(Left Shift) This operator shifts the bits of a number to the left by a specified number of positions.
									// It effectively multiplies the number by 2 raised to the power of the shift count.
	fmt.Printf("(x)08b = %08b (x)d = %d\n",x,x)			// |(Bitwise OR) This operator performs a bitwise OR operation between two numbers. 
									// It sets each bit of the result to 1 if at least one corresponding bit of the operands is 1.  
	fmt.Printf("%08b\n",y)											
	fmt.Printf("(x&y)08b = %08b (x&y)d = %d\n",x&y, x&y)		// &(Bitwise AND) This operator performs a bitwise AND operation between two numbers.
									// It sets each bit of the result to 1 if both corresponding bits of the operands are 1; otherwise, it sets the bit to 0.
	fmt.Printf("%08b\n",x|y)					// 
	fmt.Printf("%08b\n",x^y)					// ^(Bitwise XOR) This operator performs a bitwise XOR (exclusive OR) operation between two numbers.
									// It sets each bit of the result to 1 if exactly one corresponding bit of the operands is 1; otherwise, it sets the bit to 0.
	fmt.Printf("%08b\n",x&^y)

	for i := uint(0); i < 8; i++ {					
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}
	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)
}
