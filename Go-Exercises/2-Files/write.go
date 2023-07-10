package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//create blank txt file
	file, err := os.Create("file.txt")
	if err != nil {
		return
	}
	defer file.Close()

	//create array of strings
	var var1 = []string{"name1", "name2", "name3", "name4", "name5"}

	//concat array of strings into single string
	output := strings.Join(var1[:], "\n")

	//create output
	file.WriteString(output)
	fmt.Println("Text file created.")

	//rename file
	os.Rename("file.txt", "newfile.txt")
	fmt.Println("File renamed.")
}
