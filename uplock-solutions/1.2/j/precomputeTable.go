package main
import "fmt"

func main () {
	var x uint64 = 18446744073709551615 			// maximun value that can be represented by a uint64
	var pc [256]byte 					// [] with 256 zeros
	fmt.Println(pc)
	for i := range pc {
		fmt.Println("pc[i] = %d",byte(pc[i]))		// pc[i] 
		fmt.Println("pc[i/2] = %d",pc[i/2])		// pc[i/2] 
		fmt.Println("byte(i&1) = %d",byte(i&1))		// byte(i&1) 
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Println("pc[i] /2 + i&1 = %d",pc[i])
	}
	 fmt.Println("end init",(pc))
	 fmt.Println("Shift right 0*8",(x>>(0*8)))
	 fmt.Println("shift right 1*8",(x>>(1*8)))
	 fmt.Println("shift right 2*8",(x>>(2*8)))
	 fmt.Println("shift right 3*8",(x>>(3*8)))
	 fmt.Println("shift right 4*8",(x>>(4*8)))
	 fmt.Println("shift right 5*8",(x>>(5*8)))
	 fmt.Println("shift right 6*8",(x>>(6*8)))
	 fmt.Println("shift right 7*8",(x>>(7*8)))
	
}
