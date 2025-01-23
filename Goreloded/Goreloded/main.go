package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	// "goreloded"
)

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filename, err)
	}
	return string(data)
}

func moveToResult(result string) {
	args2 := os.Args[2]
	err := os.WriteFile(args2, []byte(result), 0644)
	if err != nil {
		log.Fatalf("Failed to write to file %s: %v", args2, err)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Missing arguments: input file and output file are required.")
		return
	}

	// Read input file
	inputFile := args[0]
	outputFile := args[1]
	text := readFile(inputFile)

	// Process text
	result := processText(text)

	// Convert processed lines back to a single string
	finalResult := ""
	for i, line := range result {
		finalResult += strings.Join(line, " ")
		if i < len(result)-1 {
			finalResult += "\n"
		}
	}

	// Write to output file
	moveToResult(finalResult)
	fmt.Println("Processing complete. Output written to", outputFile)
}
