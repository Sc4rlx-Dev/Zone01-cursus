package main

import (
	"regexp"
	"strings"
)

func processPunctuation(words []string) []string {
	punctuations := []string{".", ",", "!", "?", ":", ";"}

	// Handle punctuation placement
	for i := 0; i < len(words); i++ {
		// If the word ends with punctuation, ensure it is attached to the previous word
		for _, p := range punctuations {
			if len(words[i]) > 1 && strings.HasSuffix(words[i], p) {
				// Handle cases like "word ,"
				if i > 0 && words[i][0] == ',' {
					words[i-1] += p
					words[i] = strings.TrimPrefix(words[i], p)
					if words[i] == "" {
						words = append(words[:i], words[i+1:]...)
						i--
					}
				}
			}
		}
	}

	// Remove unnecessary spaces before punctuation
	re := regexp.MustCompile(`\s*([.,!?;:])\s*`)
	text := strings.Join(words, " ")
	text = re.ReplaceAllString(text, "$1 ")
	text = strings.ReplaceAll(text, " .", ".") // Fix space before period
	text = strings.ReplaceAll(text, " ~", "~") // Fix space before period
	text = strings.ReplaceAll(text, " @", "@") // Fix space before period
	text = strings.ReplaceAll(text, " ,", ",") // Fix space before period
	text = strings.ReplaceAll(text, " !", "!") // Fix space before period
	text = strings.ReplaceAll(text, " ?", "?") // Fix space before period
	text = strings.ReplaceAll(text, " :", ":") // Fix space before period
	text = strings.ReplaceAll(text, " ;", ";") // Fix space before period
	text = strings.TrimSpace(text)
	return strings.Fields(text)
}
