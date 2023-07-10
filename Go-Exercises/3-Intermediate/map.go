package main

import "fmt"

func main() {
	// elements := make(map[string]string)
	// elements["O"] = "Oxygen"
	// elements["Ca"] = "Calcium"
	// elements["C"] = "Carbon"

	// fmt.Println(elements["C"])

	//simplier notation
	elements := map[string]string{
		"O":  "Oxygen",
		"Ca": "Calcium",
		"C":  "Carbon",
	}

	fmt.Println(elements["O"])

	//can do nested map of strings (similar to json format)
}
