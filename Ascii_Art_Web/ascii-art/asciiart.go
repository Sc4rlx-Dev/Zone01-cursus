package asciiartweb

import (
	"ascii-art/function"
	"fmt"
	"os"
	"strings"
)

func AsciiArt(font string, input string) (string, error) {
	currstdhash, err := function.GetHash("standard.txt")
	currshdhash, err2 := function.GetHash("shadow.txt")
	currttyhash, err3 := function.GetHash("thinkertoy.txt")
	if err != nil || err2 != nil || err3 != nil {
		return "", fmt.Errorf("failed to get file hash: %v", err)
	}
	stdhash := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	shdhash := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	ttyhash := "5192d52b75ed5bd02cf3ee4f0300bc815e7554f641181c4ef151b60cead9abb9"
	if currstdhash != stdhash || currshdhash != shdhash || currttyhash != ttyhash {
		return "", fmt.Errorf("invalid file hash, file is not the valid one")
	}

	ascii, err := os.ReadFile(font + ".txt")
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	lines := function.Split(string(ascii))
	asciiarr := make(map[rune][]string)
	chatr := 32
	for _, v := range lines {
		asciiarr[rune(chatr)] = v
		chatr++
	}

	var slice []string
	str := ""
	input = strings.ReplaceAll(input, "\r\n", "\n")
	for _, x := range input {
		if x != '\n' {
			
			str += string(x)
		} else {
			if str != "" {
				slice = append(slice, str)
				str = ""
			} else {
				slice = append(slice, "\n")
			}
		}
	}
	if str != "" {
		slice = append(slice, str)
		str = ""
	}
	var result strings.Builder
	for i := 0; i < len(slice); i++ {
		if slice[i] != "\n" {
			lastElement := function.GetArr(slice[i], asciiarr)
			
			for i := 0; i < len(lastElement); i++ {
				result.WriteString(lastElement[i] + "\n")
			}
		} else {
			result.WriteString("\n")
		}
	}
	return result.String(), nil
}
