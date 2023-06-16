package main

import "fmt"

func main() {
	//create var
	var x string = "something"

	/*create & store pointer (type has to be same as var)
	convention: var var_name *var-type*/
	var xPointer *string
	xPointer = &x

	fmt.Println("Pointer is located at: \n", xPointer)
}
