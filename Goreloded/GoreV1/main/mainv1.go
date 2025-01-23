package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) string {
	data, _ := os.ReadFile(filename)
	return string(data)
}

func moveToResult(result string) {
	args2 := os.Args[2]
	file, _ := os.OpenFile(args2, os.O_RDWR, 0o644)
	data := []byte(result)
	err := os.WriteFile(args2, data, 0o777)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func processText(givenText string) [][]string {
	allLines := strings.Split(givenText, "\n")
	result := [][]string{}
	for _, line := range allLines {
		words := strings.Fields(line)
		words = processFlags(words)
		words = processPunctuation(words)
		words = processMarks(words)
		words = processVoile(words)
		words = processFlags(words)
		result = append(result, words)
	}
	return result
}

func processFlags(words []string) []string {
	for i, word := range words {
		switch word {
		case "(cap)":
			if i == len(words)-1 && len(words) > 1 {
				if isAlpha(words[i-1]) {
					words[i-1] = capitalize(words[i-1])
					words = words[:i]
					break
				} else {
					words = words[:i]
					return processFlags(words)
				}
			} else if i != 0 {
				if isAlpha(string(words[i-1])) {
					words[i-1] = capitalize(words[i-1])
					words = append(words[:i], words[i+1:]...)
				} else if !isAlpha(string(words[i-1][0])) {
					words = append(words[:i], words[i+1:]...)
				}
				return processFlags(words)
			} else if i == 0 {
				words = words[i+1:]
				return processFlags(words)
			}
		case "(up)":
			if i == len(words)-1 && len(words) > 1 {
				words[i-1] = strings.ToUpper(words[i-1])
				words = words[:i]
				break
			} else if i == 0 {
				words = words[i+1:]
				return processFlags(words)
			} else {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
				return processFlags(words)
			}
		case "(hex)":
			if i == len(words)-1 && len(words) > 1 {
				words[i-1] = hexToInt(words[i-1])
				words = words[:i]
				break
			} else if i == 0 {
				words = words[i+1:]
				return processFlags(words)
			} else {
				words[i-1] = hexToInt(words[i-1])
				words = append(words[:i], words[i+1:]...)
				return processFlags(words)
			}
		case "(bin)":
			if i == len(words)-1 && len(words) > 1 {
				words[i-1] = binToInt(words[i-1])
				words = words[:i]
				break
			} else if i == 0 {
				words = words[i+1:]
				return processFlags(words)
			} else {
				words[i-1] = binToInt(words[i-1])
				words = append(words[:i], words[i+1:]...)
				return processFlags(words)
			}
		case "(up,":
			r, _ := strconv.Atoi(string(words[i+1][0]))
			if r > len(words[:i]) {
				r = len(words[:i])
			}
			for j := 1; j <= r; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			return processFlags(words)
		case "(cap,":
			r, _ := strconv.Atoi(string(words[i+1][:len(words[i+1])-1]))
			if r > len(words[:i]) {
				r = len(words[:i])
			}
			for j := 1; j <= r; j++ {
				words[i-j] = capitalize(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			return processFlags(words)
		case "(low)", "(low,":
			if word == "(low)" {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
				return processFlags(words)
			} else {
				r, _ := strconv.Atoi((string(words[i+1][0])))
				if r > len(words[:i]) {
					r = len(words[:i])
				}
				for j := 1; j <= r; j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
				return processFlags(words)
			}
		}
	}
	return words
}

func capitalize(s string) string {
	index := 0
	s2 := strings.ToLower(s)
	for i := range s2 {
		if isAlpha(string(s2[i])) {
			index = i
			break
		}
	}
	x := s2[index] - 32
	first := s2[:index]
	last := s2[index+1:]
	s3 := first + string(x) + last

	return s3
}

func processPunctuation(words []string) []string {
	punctuations := []byte{'.', ',', '!', '?', ':', ';'}
	for i, word := range words {
		ln := len(word)
		for ln >= 1 {
			for _, punctuation := range punctuations {
				if word[0] == punctuation && ln > 1 && i != 0 {
					words[i] = word[1:]
					words[i-1] = words[i-1] + string(punctuation)
					return processPunctuation(words)
				} else if word[0] == punctuation && ln == 1 && i != 0 {
					words[i-1] = words[i-1] + string(punctuation)
					words = append(words[:i], words[i+1:]...)
					return processPunctuation(words)
				}
			}
			ln--
		}
	}
	return words
}

func isAlpha(s string) bool {
	myString := []rune(s)
	for _, char := range myString {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			return true
		}
	}
	return false
}

func checkWord(words []string, mark byte) int {
	count := 0
	for _, word := range words {
		for _, char := range word {
			if byte(char) == mark {
				if char == rune(mark) && (char == rune(word[0]) || char == rune(word[len(word)-1])) {
					count++
				}
			}
		}
	}
	return count
}

func convertToString(words []string) string {
	str := ""
	for i, word := range words {
		if i != 0 {
			str += " " + word
		} else {
			str += word
		}
	}
	return str
}

func processMarks(words []string) []string {
	mark := byte(39)
	for i, word := range words {
		lnw := len(word)
		for c, char := range word {
			if char == rune(mark) && len(words) == 1 {
				break
			} else if char == rune(mark) && lnw == 1 {
				if (i == 0 || i == len(words)-1) && len(words) > 1 {
					if i == 0 && words[i+1][0] != mark {
						words[i+1] = string(mark) + words[i+1]
						words = words[i+1:]
						return processMarks(words)
					} else if checkWord(words[:i], mark)%2 != 0 {
						words[i-1] = words[i-1] + string(mark)
						words = words[:i]
						return processMarks(words)
					}
				} else {
					if checkWord(words[:i], mark)%2 != 0 {
						words[i-1] = words[i-1] + string(mark)
						words = append(words[:i], words[i+1:]...)
						return processMarks(words)
					} else if words[i+1][0] != mark {
						words[i+1] = string(mark) + words[i+1]
						words = append(words[:i], words[i+1:]...)
						return processMarks(words)
					}
				}
			} else if char == rune(mark) && lnw > 1 {
				if c == 0 && i != len(words)-1 {
					if checkWord(words[:i], mark)%2 != 0 && (words[i-1][len(words[i-1])-1] != mark) {
						words[i-1] = words[i-1] + string(mark)
						words[i] = words[i][1:]
					}
				} else if c == 0 && i == len(words)-1 {
					if checkWord(words[:i], mark)%2 != 0 && ((words[i-1][len(words[i-1])-1] != mark) || len(words[i-1]) == 1) {
						words[i-1] = words[i-1] + string(mark)
						words[i] = words[i][1:]
					}
				} else if c == len(word)-1 {
					if i == len(words)-1 {
						break
					} else if words[i+1][0] != mark {
						words[i+1] = string(mark) + words[i+1]
						words[i] = words[i][:c]
					}
				}
			}
		}
	}
	return words
}

func processVoile(words []string) []string {
	for i, word := range words {
		if i != 0 {
			if words[i-1] == "a" || words[i-1] == "A" {
				voile := []byte{'a', 'e', 'i', 'o', 'u', 'h'}
				for _, v := range voile {
					if word[0] == v || word[0] == v-32 {
						words[i-1] = words[i-1] + "n"
						break
					}
				}
			}
		}
	}
	return words
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("missing argument")
	} else if len(args) == 2 {
		str := ""
		result := processText(readFile(args[0]))
		for s, slice := range result {
			if len(result) > 1 && s < len(result)-1 {
				str += convertToString(slice) + "\n"
			} else {
				str += convertToString(slice)
			}
		}
		moveToResult(str)
		fmt.Println("success!")
	}
}

func hexToInt(hex string) string {
	hexConvert, err := strconv.ParseInt(hex, 16, 64)
	hex2 := ""
	if err != nil {
		hex2 = hex
	} else {
		hex2 = strconv.FormatInt(hexConvert, 10)
	}
	return hex2
}

func binToInt(bin string) string {
	binConvert, err := strconv.ParseInt(bin, 2, 64)
	bin2 := ""
	if err != nil {
		bin2 = bin
	} else {
		bin2 = strconv.FormatInt(binConvert, 10)
	}
	return bin2
}
