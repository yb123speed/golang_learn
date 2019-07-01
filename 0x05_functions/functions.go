package main

import (
	"fmt"
)

func main(){
	fmt.Println("Total price is ", calculateBill(10, 20))

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f \r\n", area, perimeter) 
	
	area1, _ := rectProps(10.8, 5.6) // perimeter is discarded
    fmt.Printf("Area %f ", area1)
}

func calculateBill(price int, no int) int{
	var totalPrice = price * no
    return totalPrice
}

func rectProps(length, width float64) (float64, float64){
	return length*width, 2*(length+width)
}