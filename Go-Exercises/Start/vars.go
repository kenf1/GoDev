package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	/*calculate yr given DoB & age
	user inputs age, print yr they were born*/

	//prompt user for age
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter age:")
	age, _ := reader.ReadString('\n')
	age1 := strings.Replace(age, "\n", "", -1)

	//convert age1 to int
	age_num, e := strconv.Atoi(age1)
	if e != nil {
		fmt.Println("Entered value is not an accepted value.")
	} else {
		//get current date (don't need month & day)
		yr, _, _ := time.Now().Date()
		// fmt.Println(yr)

		birth_yr := yr - age_num
		//sep into 2 lines to avoid converting birth_yr back to string
		fmt.Println("Based on the age you entered, you were born on:")
		fmt.Println(birth_yr)
	}
}
