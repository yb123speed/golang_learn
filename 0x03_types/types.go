package main

import "fmt"

func main(){

    // Complex
    var a, b = 3.1415, 1.234
    cplx := complex(a, b)
    fmt.Printf("cplx is %T \n", cplx)

    cplx1 := complex(7, 5)
    fmt.Printf("cplx1 is %T \n", cplx1)
    fmt.Println("cpx1 is ", cplx1)

    // Type Conversion

    // i := 55      //int
    // j := 67.8    //float64
    // sum := i + j //Error invalid operation: i + j (mismatched types int and float64)
    // fmt.Println(sum)

    i := 55      //int
    j := 67.8    //float64
    sum := i + int(j) //int
    fmt.Println(sum)

    i1 := 10
    var j1 float64 = float64(i1) //this statement will not work without explicit conversion
    fmt.Println("j1", j1)
}