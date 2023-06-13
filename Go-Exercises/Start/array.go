package main

import "fmt"

func main() {
	ex_array := []int64{1, 2, 3}

	/*for i in 1:length(ex_arry){
		print(ex_array[i])
	}*/
	for i := 0; i < len(ex_array); i++ {
		fmt.Println(ex_array[i])
	}

	fmt.Println("---")

	//alternative
	for _, num := range ex_array {
		fmt.Println(num)
	}
}
