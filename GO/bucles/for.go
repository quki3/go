package main

import "fmt"

func main(){
	//this is a for normal
	i := 0;
	for i<10 {
		fmt.Println(i)
		i++
	}
	//break an cicle infinite
	j :=0;
	for {
		if i>9{
			break
		}
		fmt.Println(i)
		i++
	}

	// modulo % da el resto de una divicion
	x:=1

	for {
		x++
		if x>100{
			breack
		}
		if x%2!=0{//if x es distinto de impar va a continuar es decir va a interar nuevamente
			continue
		}
		fmt.Println(x)
	}

	//how print values acii whit for
	for i:=33; i<=122; i++{
		fmt.Printf("%d\t%#x\t%#U\n",i,i,i)
}









