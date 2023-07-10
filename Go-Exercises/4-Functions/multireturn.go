package main

import "fmt"

func values() (string, int) {
	//change values to hello world (string)
	// return "hello", "world"

	//mix of diff types
	return "hello", 1
}

func main() {
	x, y := values()
	fmt.Println(x)
	fmt.Println(y)
}
