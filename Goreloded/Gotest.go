package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// line := scanner.Text()
		// process the line
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	data , err := os,ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	// data
	
}
