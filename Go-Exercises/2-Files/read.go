package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func linebyline(path string) ([]string, error) {
	//read in file line by line
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//create slice to store lines as array
	var lines []string

	//open file & append each line to array
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	//read in names.txt
	lines, err := linebyline("names.txt")
	if err != nil {
		log.Fatalf("linebyline: %s", err)
	}

	//print index followed by line content
	for i, line := range lines {
		fmt.Println(i, line)
	}
}
