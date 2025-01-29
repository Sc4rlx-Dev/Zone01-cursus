package main

import "strings"

func capitalize(s string) string {
	if len(s) == 0 {
		return s }
	runes := []rune(strings.ToLower(s))
	runes[0] = rune(strings.ToUpper(string(runes[0]))[0])
return string(runes)}
