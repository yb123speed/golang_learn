package main

import "fmt"

func main(){
    const a = 100
    //a = 99 //not allowed
    fmt.Println("a is ", a)

    const b = "Hello World"
    fmt.Printf("the type of b is %T \n", b)

    var name = "Sam"
    fmt.Printf("type %T value %v\n", name, name)

    // number
    const a1 = 5
    var intVar int = a1
    var int32Var int32 = a1
    var float64Var float64 = a1
    var complex64Var complex64 = a1
    fmt.Printf("the type of a1 is %T, the value is %v \n", a1, a1)
    fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)

    var a2 = 5.9/8
    fmt.Printf("a2's type %T value %v \n",a2 , a2)
}