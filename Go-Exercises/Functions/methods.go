package main

import (
	"fmt"
)

func sum(a, b float64) float64 {
	return a + b
}

func main() {
	funcResult := sum(6.1, 9.3)
	fmt.Println("The answer is: " + fmt.Sprint(funcResult))

	//defer
	defer fmt.Println("Hello")
	defer fmt.Println("!")
	fmt.Println("World")

	/*prints out World ! Hello
	world doesn't have defer -> runs 1st
	! is last defer run -> goes next
	finally Hello
	*/
}
