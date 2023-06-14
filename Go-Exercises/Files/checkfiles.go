package main

import (
	"fmt"
	"os"
)

//checks if sample.txt exists; throws error if not

func main() {
	if _, err := os.Stat("sample.txt"); os.IsNotExist(err) {
		fmt.Println("File doesn't exist.")
	} else {
		fmt.Println("File exists.")
	}
}
