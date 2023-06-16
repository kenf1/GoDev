package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//use current time as seed
	rand.Seed(time.Now().UnixNano())

	//select random # from 0 -> 5, then shift frame to be 1 -> 6
	diceNum := rand.Intn(6) + 1

	fmt.Println("Dice rolled: \n", diceNum)
}
