# GO.__GB__
## Ladies and gentlemen this is GO
#  (imprimir hola mundo)
```bash

package main
import "fmt"

func main(){
    fmt.Println("hola Mundo")
}
```
# (operador de declaración corta)
```bash
package main
import "fmt"

func main(){
    x:= 42
    y:= "James Bond"
    z:=true
    fmt.Println(x,y,z)
    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)
}
```

# (las variables son de los siguientes tipos)
```bash
package main
import "fmt"

var x int
var y string
var z bool

func main(){
    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)
}
```
# ejercicio 4 (tipo subyacente, raíz o tipo implícito %T)
```bash
package main 
import("fmt")

type numero int
var x numero
func main(){
            fmt.println(x)
            fmtprintf("El tipo de x es : %T\n", x)
            
            x=42 //? le asigno a x 42
            fmt.println(x)
            
            y = int(x)
            fmt.Println(y)
            fmt.Printf("%T",y)
}
```
`fmt.printf` especifica el formato
`%T` para impromir Tipos es decir de que tipo es el valor
`\n` inprime una linea nueva
# ejercicio #5 
```bash
package main
import ("fmt")

type numero int
var x numero
var y int
func main{
    fmt.println(x)
    fmt.printf('el tipo de x es: %T, x)
    x = 20
    fmt.Println(x)
    
    y =int(x)
    fmt.Println(y)
    fmt.Printf("%T",y)
    
}
```


