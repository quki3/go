//package main
package main
//import fmt
import "fmt"
//declaramos cosnt

//creamos la func main
func main(){
	//i init i<= condition i++ post
	for i:=0; i<=100; i++{
		fmt.Println(i)
	}
}
//-------------------------------------------
package main
import "fmt"
//const

//func main
func main(){
	for i:=0; i<=10;i++{
		fmt.Printf("El ciclo externo %d\n",i)
		for j:=0; j<=10;j++{
			fmt.Printf("El ciclo interno %d\n",j)
		}
	}	
}
