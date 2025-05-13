package function

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func Split(ascii string) [][]string {
	slise := []string{}
	lines := [][]string{}
	line := ""
	for _, v := range ascii {
		if v == '\n' {
			slise = append(slise, line)
			line = ""

		} else {
			line += string(v)
		}
	}
	for i := 1; i < len(slise); i += 8 {
		if i+8 > len(slise) {
			break
		}
		abed := slise[i : i+8]
		lines = append(lines, abed)
		i++
	}
	return lines
}

func GetArr(input string, asciiarr map[rune][]string) []string {
	ret := make([]string, 8)

	for _, v := range input {
		ris := asciiarr[v]
		for i := 0; i < 8; i++ {
			if i < len(ris) {
				ret[i] += ris[i]
			}
		}
	}
	return ret
}

func GetHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha256.New()

	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
