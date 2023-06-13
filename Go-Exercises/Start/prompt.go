package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//prompt user for name
	// fmt.Println("Enter your name: ")
	// name, _ := reader.ReadString('\n')
	// fmt.Println("The name you entered is: " + name)

	//prompt user for # & check if 1 ≤ x ≤ 10
	fmt.Println("Enter a number: ")
	str, _ := reader.ReadString('\n')

	//rm \n from string
	str1 := strings.Replace(str, "\n", "", -1)

	/*convert string to int
	num check ran only if conversion success*/
	num, e := strconv.Atoi(str1)
	if e != nil {
		fmt.Println(str1 + " is not an accepted value.")
	} else {
		//check if 1 ≤ x ≤ 10
		if num >= 1 && num <= 10 {
			fmt.Println("Entered number is between 1 and 10.")
		} else {
			fmt.Println("Entered number is NOT between 1 and 10.")
		}
	}
}
