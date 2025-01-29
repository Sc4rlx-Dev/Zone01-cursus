package main

import ("strings")

func processPunctuation(words []string) []string {
	pct := []string{".", ",", "!", "?", ":", ";"}
	res := []string{} 
	finalres := []string{} 

	for i, word := range words {
		for _, p := range pct {
			if strings.Contains(word, p) {	word = strings.ReplaceAll(word, " "+p, p)
				if strings.HasSuffix(word, p) && i < len(words)-1 { word = strings.TrimSuffix(word, p) + p }
			}
		}
		res = append(res, word)}
	for i, word := range res {
		if i > 0 && contains(pct, string(word[0])) { finalres[len(finalres)-1] += word
		} else {   finalres = append(finalres, word)  }
	}
	return finalres}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str { return true }
	}
return false}
