package main
import (
	"fmt"
	"maps"
)
func main(){
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:",m)
	v1 := m["k1"]
	fmt.Println("v1:",v1)
	v3 := m["k3"]
	fmt.Println("v3:",v3)
