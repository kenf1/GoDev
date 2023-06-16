package main

import (
	"fmt"
	"strings"
)

type words struct {
	w1 string
	w2 string
}

func main() {
	sliceEx := "hello world"

	var splitEx words
	splitEx.w1 = sliceEx[0:5]
	splitEx.w2 = sliceEx[6:11]

	fmt.Println("The string:", sliceEx, "is split by position into:\n", splitEx.w1, "\n", splitEx.w2)

	//alternatively
	altEx := strings.Split(sliceEx, " ")
	fmt.Println("Alternative way is strings.Split:\n", altEx[0], "\n", altEx[1])
}
