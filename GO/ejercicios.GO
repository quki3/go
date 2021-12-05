// GO.__GB__
// Ladies and gentlemen this is GO
//  (imprimir hola mundo)


package main
import "fmt"

func main(){
    fmt.Println("hola Mundo")
}

// (operador de declaración corta)

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


// (las variables son de los siguientes tipos)

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

// ejercicio 4 (tipo subyacente, raíz o tipo implícito %T)

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

//`fmt.printf` especifica el formato
//`%T` para impromir Tipos es decir de que tipo es el valor
//`\n` inprime una linea nueva

// ejercicio #5 

package main
import ("fmt")

type numero int
var x numero
var y int
func main{
    fmt.println(x)
    fmt.printf('el tipo de x es: %T', x)
    x = 20
    fmt.Println(x)
    
    y =int(x)
    fmt.Println(y)
    fmt.Printf("%T",y)
    
}

//EJERCICIO NRO 6 

package main
import("fmt")
func main(){
    fmt.Printf("%d\t%b\t%x",a,a,a)
}

/*Ejercicio práctico #2
Usando los siguientes operadores, escribe expresiones y asigna sus valores a variables:
==
<=
>=
!=
<
>
Imprime los valores de las variables. 
*/
package main

import (
    "fmt"
)
func main(){
    //declaramos y asignamos lo que nos retorna la expresion
    a := (42 == 42)
    b := (42 <= 42)
    c := (42 >= 42)
    d := (42 != 41)
    e := (42 < 43)
    f := (42 > 41)
    //imprimir por consola
    fmt.println(a,b,c,d,e,f)
}

/*Ejercicio práctico #3
Crea constantes CON TIPO y SIN TIPO. Imprime el valor de las mismas.*/

//definimos el paclkage
package main
//importamos los modulos que usamos
import "fmt"

//creamos constantes
cosnt (
    a = 42
    b int = 43
)

//definimos la funcion a usar
func main(){
    fmt.println(a,b)
}

/*Ejercicio práctico #4
Escribe un programa que 
Asignar un int a una variable
Imprímelo en decimal, binario, y hex
Has shifts de bits de ese int una posición a la izquierda y asigna eso a una variable
Imprime esa variable en decimal, binario, y hex
*/
//definimos package
package main
//import fmt
import ("fmt")
