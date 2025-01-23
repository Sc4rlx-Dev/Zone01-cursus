package main

import (
	"strings"
)

func processText(givenText string) [][]string {
	lines := strings.Split(givenText, "\n")
	var result [][]string
	for _, line := range lines {
		words := strings.Fields(line)
		words = processFlags(words)
		words = processPunctuation(words)
		words = processMarks(words)
		words = processVoile(words)
		words = processFlags(words) // Ensure all flags are handled
		result = append(result, words)
	}
	return result
}
