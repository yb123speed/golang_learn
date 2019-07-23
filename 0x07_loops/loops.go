package main

import "fmt"

func main(){
    for i:=1;i<=10;i++ {
        fmt.Printf(" %d", i)
    }

    fmt.Println()

    for i:=1;i<=10;i++ {
        if i>5 {
            break;
        }
        fmt.Printf(" %d", i)
    }
    fmt.Printf("\nline after for loop")

    fmt.Println()

    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Printf("%d ", i)
    }

    fmt.Println()

    n := 5
    for i := 0; i < n; i++ {
        for j := 0; j <= i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }

    outer:  
    for i := 0; i < 3; i++ {
        for j := 1; j < 4; j++ {
            fmt.Printf("i = %d , j = %d\n", i, j)
            if i == j {
                break outer
            }
        }
    }
}