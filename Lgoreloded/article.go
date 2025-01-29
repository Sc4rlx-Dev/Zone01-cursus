package main

import ( "strings" )

func processArticle(w []string) []string {
	vowels := "aeiouhAEIOUH"
	for i := 1; i < len(w); i++ {
		if (w[i-1] == "a" || w[i-1] == "A") && strings.ContainsRune(vowels, rune(w[i][0])) { w[i-1] += "n" }}
return w }
