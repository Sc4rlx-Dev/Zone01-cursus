package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	// tari9a "go-reloaded"
)

// this  fun used  to read  content of  input  file
func ReadFile(filename string) string {
	data, _ := os.ReadFile(filename)
	return string(data)
}

// this  func  to past  result  in  result file
func MovetoResult(result string) {
	args2 := os.Args[2]
	file, _ := os.OpenFile(args2, os.O_RDWR, 0o644)
	data := []byte(result)
	err := os.WriteFile(args2, data, 0o777)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

// this  func  used  to caall  of  all  func i need
func correction(giventext string) [][]string {
	allline := strings.Split(giventext, "\n")
	result := [][]string{}
	for j := 0; j < len(allline); j++ {
		words := strings.Fields(allline[j])
		words = flags(words)
		words = poctuation(words)
		words = marks(words)
		words = voile(words)
		words = flags(words)
		result = append(result, words)
	}
	return result
}

// this  func  used  for  flags
func flags(slice []string) []string {
	words := slice
	for i, word := range words {
		if word == "(cap)" {
			if i == len(words)-1 && len(words) > 1 {
				if IsAlpha(words[i-1]) {
					words[i-1] = tocap(words[i-1])
					words = words[:i]
					break
				} else {
					words = words[:i]
					return flags(words)
				}
			} else if i != 0 {
				if IsAlpha(string(words[i-1])) {
					words[i-1] = tocap(words[i-1])
					words = append(words[:i], words[i+1:]...)
				} else if !IsAlpha(string(words[i-1][0])) {
					words = append(words[:i], words[i+1:]...)
				}
				return flags(words)
			} else if i == 0 {
				words = words[i+1:]
				return flags(words)
			}
		} else if word == "(up)" {
			if i == len(words)-1 && len(words) > 1 {
				words[i-1] = strings.ToUpper(words[i-1])
				words = words[:i]
				break
			} else if i == 0 {
				words = words[i+1:]
				return flags(words)
			} else {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
				return flags(words)
			}
		} else if word == "(hex)" {
			if i == len(words)-1 && len(words) > 1 {
				words[i-1] = Hextoint(words[i-1])
				words = words[:i]
				break
			} else if i == 0 {
				words = words[i+1:]
				return flags(words)
			} else {
				words[i-1] = Hextoint(words[i-1])
				words = append(words[:i], words[i+1:]...)
				return flags(words)
			}
		} else if word == "(bin)" {
			if i == len(words)-1 && len(words) > 1 {
				words[i-1] = Bintoint(words[i-1])
				words = words[:i]
				break
			} else if i == 0 {
				words = words[i+1:]
				return flags(words)
			} else {
				words[i-1] = Bintoint(words[i-1])
				words = append(words[:i], words[i+1:]...)
				return flags(words)
			}
		} else if word == "(up," {
			r9em0, _ := strconv.Atoi(string(words[i+1][0])) // convert   from  byte to int
			if r9em0 > len(words[:i]) {
				r9em0 = len(words[:i])
			}
			for j := 1; j <= r9em0; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			return flags(words)
		} else if word == "(cap," {
			r9em, _ := strconv.Atoi(string(words[i+1][:len(words[i+1])-1]))
			if r9em > len(words[:i]) {
				r9em = len(words[:i])
			}
			for j := 1; j <= r9em; j++ {
				words[i-j] = tocap(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			return flags(words)
		} else if word == "(low)" || word == "(low," {
			if word == "(low)" {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
				return flags(words)
			} else {
				r9em, _ := strconv.Atoi((string(words[i+1][0])))
				if r9em > len(words[:i]) {
					r9em = len(words[:i])
				}
				for j := 1; j <= r9em; j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
				return flags(words)
			}
		}
	}
	return words
}

// to  capitalize word
func tocap(s string) string {
	index := 0
	s2 := strings.ToLower(s)
	for i := range s2 {
		if IsAlpha(string(s2[i])) {
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

// this  function use  for procces poctuation
func poctuation(slice []string) []string {
	punctuations := []byte{'.', ',', '!', '?', ':', ';'}
	s := slice
	count := 0
	for i := range s {
		ln := len(s[i])
		for ln >= 1 {
			for j := range punctuations {
				if s[i][count] == punctuations[j] && ln > 1 && i != 0 {
					s[i] = s[i][1:]
					s[i-1] = s[i-1] + string(punctuations[j])
					return poctuation(s)
				} else if s[i][count] == punctuations[j] && ln == 1 && i != 0 {
					s[i-1] = s[i-1] + string(punctuations[j])
					s = append(slice[:i], slice[i+1:]...)
					return poctuation(s)
				}
			}
			ln--
		}
	}
	return s
}

// this  for check  the  first  char  in word  is  an  alpha  used just  in capitalize function
func IsAlpha(s string) bool {
	mystring := []rune(s)
	a := false
	for i := 0; i < len(mystring); i++ {
		if mystring[i] >= 'a' && mystring[i] <= 'z' || (mystring[i] >= 'A' && mystring[i] <= 'Z') {
			a = true
			break
		}
	}
	return a
}

// thes  for  check the  marks in  the  last  argumment
func checkword(slice []string, mark byte) int {
	count := 0
	for w, word := range slice {
		for x := range slice[w] {
			if slice[w][x] == mark {
				if x == 0 || x == len(word)-1 {
					count++
				}
			}
		}
	}
	return count
}

// thes  for  convert  a slice to  a string
func converts(slice []string) string {
	// s := slice
	str := ""
	for w, word := range slice {
		if w != 0 {
			str = str + " " + word
		} else {
			str = word
		}
	}
	return str
}

// this  for marks '
func marks(slice2 []string) []string {
	// var  s2  []string
	s := slice2
	var mark byte = 39
	for w, word := range s {
		lnw := len(word)
		for c, char := range word {
			if char == rune(mark) && len(s) == 1 {
				break
				// if ' ==  word
			} else if char == rune(mark) && lnw == 1 {
				// if  the mark  is  in the first  of text  or last
				if (w == 0 || w == len(s)-1) && len(s) > 1 {
					if w == 0 && s[w+1][0] != mark {
						s[w+1] = string(mark) + s[w+1]
						s = s[w+1:]
						return marks(s)
					} else if checkword(s[:w], mark)%2 != 0 {
						s[w-1] = s[w-1] + string(mark)
						s = s[:w]
						return marks(s)
					}
					// if   ' in the botom  of  text
				} else {
					if checkword(s[:w], mark)%2 != 0 {
						s[w-1] = s[w-1] + string(mark)
						s = append(s[:w], s[w+1:]...)
						return marks(s)
					} else if s[w+1][0] != mark {
						s[w+1] = string(mark) + s[w+1]
						s = append(s[:w], s[w+1:]...)
						return marks(s)
					}
				}
				// if  '  is  just  a char
			} else if char == rune(mark) && lnw > 1 {
				// if ' in the first  char  and  not  is  in the  first  word
				if c == 0 && w != len(s)-1 {
					if checkword(s[:w], mark)%2 != 0 && (s[w-1][len(s[w-1])-1] != mark) {
						s[w-1] = s[w-1] + string(mark)
						s[w] = s[w][1:]
					}
					//  if  '  in  the  last  of word
				} else if c == 0 && w == len(s)-1 {
					if checkword(s[:w], mark)%2 != 0 && ((s[w-1][len(s[w-1])-1] != mark) || len(s[w-1]) == 1) {
						s[w-1] = s[w-1] + string(mark)
						s[w] = s[w][1:]
					}
				} else if c == len(word)-1 {
					if w == len(s)-1 {
						break
					} else if s[w+1][0] != mark {
						s[w+1] = string(mark) + s[w+1]
						s[w] = s[w][:c]
						// return marks(s)
					}
				}
			}
		}
	}
	return s
}

// this  for  voile
func voile(slice []string) []string {
	s := slice
	voile := []byte{'a', 'e', 'i', 'o', 'u', 'h'}
	for w, word := range s {
		if w != 0 {
			if s[w-1] == "a" || s[w-1] == "A" {
				for v := range voile {
					if word[0] == voile[v] || word[0] == voile[v]-32 {
						s[w-1] = s[w-1] + "n"
						break
					}
				}
			}
		}
	}
	return s
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("missing  argumment")
	} else if len(args) == 2 {
		str := ""
		result := correction(ReadFile(args[0]))
		for s, slice := range result {
			if len(result) > 1 && s < len(result)-1 {
				str += converts(slice) + "\n"
			} else {
				str += converts(slice)
			}
		}
		MovetoResult(str)
		fmt.Println("succes !! ")
	}
}

// / this  used  for  convert  A HEX  TO int
func Hextoint(hex string) string {
	hexconvert, err := strconv.ParseInt(hex, 16, 64)
	hex2 := ""
	if err != nil {
		hex2 = hex
	} else {
		hex2 = strconv.FormatInt(hexconvert, 10)
	}
	return hex2
}

// this  used  for  covert  binary  to  number
func Bintoint(bin string) string {
	binconvert, err := strconv.ParseInt(bin, 2, 64)
	bin2 := ""
	if err != nil {
		bin2 = bin
	} else {
		bin2 = strconv.FormatInt(binconvert, 10)
	}
	return bin2
}
