package main

import (
	"strings")

func processMarks(words []string) []string {
	mark := byte(39) // ASCII for single quote '

	countMarks := func(slice []string) int {
		count := 0
		for _, word := range slice {
			count += strings.Count(word, string(mark))
		}
		return count
	}

	for i := 0; i < len(words); {
		if words[i] == string(mark) {
			if len(words) == 1 {
				break
			}

			if i == 0 && len(words) > 1 && !strings.HasPrefix(words[i+1], string(mark)) {
				words[i+1] = string(mark) + words[i+1]
				words = words[1:]
			} else if i == len(words)-1 && countMarks(words[:i])%2 != 0 {
				words[i-1] += string(mark)
				words = words[:i]
			} else if i > 0 && i < len(words)-1 {
				if countMarks(words[:i])%2 != 0 {
					words[i-1] += string(mark)
				} else if !strings.HasPrefix(words[i+1], string(mark)) {
					words[i+1] = string(mark) + words[i+1]
				}
				words = append(words[:i], words[i+1:]...)
			} else {
				i++
			}
			continue
		}

		if strings.HasPrefix(words[i], string(mark)) && len(words[i]) > 1 {
			if countMarks(words[:i])%2 != 0 {
				words[i-1] += string(mark)
				words[i] = words[i][1:]
			}
		}
		if strings.HasSuffix(words[i], string(mark)) && len(words[i]) > 1 {
			if i < len(words)-1 && !strings.HasPrefix(words[i+1], string(mark)) {
				words[i+1] = string(mark) + words[i+1]
				words[i] = words[i][:len(words[i])-1]
			}
		}
		i++
	}

	for i := range words {
		if len(words[i]) > 1 && words[i][0] == mark && words[i][len(words[i])-1] == mark {
			words[i] = strings.TrimSpace(words[i])
		}
	}

	return words
}
