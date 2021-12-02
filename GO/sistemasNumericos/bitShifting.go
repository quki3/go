/*
Bit shifting
	Bit shifting es cuando desplazan dígitos binarios a la izquierda o a la
 	derecha. Podemos usar bit shifting junto con iota, para construir constantes
  	creativas.
*/
package main

import (
	"fmt"
)

//para usar bit shifting
const (
	//desechamos el primer 0 de iota
	_ = iota

	kbBS = 1 << (iota * 10)
	bgBS = 1 << (iota * 10)
	tbBS = 1 << (iota * 10)
)

func main(){
	//ejem 1
	a:=4
	fmt.Println("%d\t\t%b",a,a)
	// <<1 Bit shifting corre el balor significativo a la izquierda
	b =a << 1
	//%d muestra un valor en decimal %b muestra un valor en binario
	fmt.Println("%d\t\t%b",b,b)

	//ejem 2 
	kb := 1024
	gb := kb * 1024
	tb := gb * 1024

	fmt.Printf("%d\t\t\t%b\n",kb,kb)
	fmt.Printf("%d\t\t\t%b\n",gb,gb)
	fmt.Printf("%d\t\t\t%b\n",tb,tb)
	
	//usando bit shifting 
	fmt.Printf("%d\t\t\t%b\n",kbBS,kbBS)
	fmt.Printf("%d\t\t\t%b\n",gbBS,gbBS)
	fmt.Printf("%d\t\t\t%b\n",tbBS,tbBS)
}