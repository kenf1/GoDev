package main

import "fmt"

type House struct {
	//struct is similar to class in OOP languages but Golang is not an OOP language
	numRooms int32
	price    float32
	city     string
}

func main() {
	var location1 House

	location1.numRooms = 3
	location1.price = 1000.50
	location1.city = "Planet Earth"

	//print single aspect of struct
	// fmt.Println("location1.city: \n", location1.city)

	//print entire struct
	fmt.Println("Location1's details:")
	fmt.Printf("%+v \n", location1)
}
