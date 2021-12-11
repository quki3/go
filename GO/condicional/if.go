package main
import ("fmt")

func main(){
	if true {
		fmt.Println(001)
	}
	if false {
		fmt.Println(002)
	}
	if !true {
		fmt.Println(003)
	}
	if !false {
		fmt.Println(004)
	}
	if 43 =43 {
		fmt.Println(005)
	}
	/*el scope de x es solo en ese bloque de codigo*/
	if x:=42; x == 42{
		fmt.Println("001")
	}
}
