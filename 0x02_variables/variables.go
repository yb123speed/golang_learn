package main

import "fmt"

func main(){
	var age int
	fmt.Println("my age is",age)
	age = 27
	fmt.Println("my age is",age)
	age = 54
	fmt.Println("my age is",age)

	var age1 int = 29
	fmt.Println("age1 is", age1)

	var age2 = 29
	fmt.Println("age2 is", age2)

	//var age3, age4 int = 33, 34
	var age3, age4 = 33, 34
	fmt.Println("age3 is", age3," age4 is",age4)

	var (
		name	= "Yebin"
		age5		= 29
		height int
	)
	fmt.Println("name is ",name, " age is ", age5, "height is", height)

	name1, age6 := "naveen", 29 //short hand declaration
    fmt.Println("my name is", name1, "age is", age6)

	name1, age7 := "yebin", 34 // one new variable
	fmt.Println("my name is", name1, "age is", age7)
}

