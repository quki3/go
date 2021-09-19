# GO.__GB__

## tipo de valores numericos

```go
//ruta /language specification/types/numeric types
int //? enteros
float64 //? numeros de puntos flotantes, numeros con comas
uint8 //? solo podra obtener valores del 0 al 255 sin incluir numeros negativos por la u que le pusimos adelante
```
## string
```go


```




## informacion de valores
```go
%T //? esto me da el tipo de dato
\n //?linea nuevas
%#U //te devuelve el unicode point uf8 de cada letra
```

## sistemas numericos
convercion de base decimal a otra base
```bash
numero decimal | base
(resto 1)        (resultado)| base
                  (resto 2)    (resultado) | base
                                  (resto 3)    (resultado final)
resultado final + resto 3 + resto 2 + resto 1 = numero.
                                                      .base
```
## constantes
```go
    package main
    import(
        "fmt"
    )
    const a = 42
    const ( 
        b = 42.32
        c = "Eduar Tua"
    )

    type nombre string
    var d nombre

    func main(){
        fmt.Println(a)
        fmt.Println(b)
        fmt.Println(c)
        fmt.Println("%T\n",a)
        fmt.Println("%T\n",b)
        fmt.Println("%T\n",c)
    }
    
```
## Iota
Dentro de una declaracion, el identificador pre-declarado iota representa una sucesion de constantes enteras sin tipo, se reinicia a 0 cada vez que se llame a const

```go
    package main
    import(
        "fmt"
    )
    const (
        a = iota 
        b = iota 
        c = iota 
    )
    const (
        d = iota //ouput 0
        f       //ouput 1
        g       //ouput 2
    )
    func main(){
        fmt.Println(a)//ouput  0
        fmt.Println(b)//ouput  1
        fmt.Println(c)//ouput  2
        fmt.Println(d)//ouput  0
        fmt.Println(f)//ouput  1
        fmt.Println(g)//ouput  2
        fmt.Printf("%T\n",a)// ouput int
        fmt.Printf("%T\n",b)//ouput int
        fmt.Printf("%T\n",c)// ouput int
        fmt.Printf("%T\n",d)// ouput int
        fmt.Printf("%T\n",f)// ouput int
        fmt.Printf("%T\n",g)//ouput int
    }
    
```
