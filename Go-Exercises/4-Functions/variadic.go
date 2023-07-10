package main

import "fmt"

func concatPrint(names ...string) {
	//similar to Python args & kwargs
	for _, name := range names {
		fmt.Println(name)
	}
}

func main() {
	concatPrint("name1", "name2", "name3")
}
