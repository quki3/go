# GO.__GB__
## Ladies and gentlemen this is GO
##  (imprimir hola mundo)
```bash

package main
import "fmt"

func main(){
    fmt.Println("hola Mundo")
}
```
## (operador de declaración corta)
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

## (las variables son de los siguientes tipos)
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
## (tipo subyacente, raíz o tipo implícito %T)
```bash
package main 
import("fmt")

type numero int
var x numero
func main(){
            fmt.println(x)
            fmtprintf("El tipo de x es : %T\n", x)
            
            x=42
            fmt.println(x)
            
            y = int(x)
            fmt.Println(y)
            fmt.Printf("%T",y)
}
```


