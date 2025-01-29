package main

import ("strings")

func processText(txt string) [][]string {
	lines := strings.Split(txt, "\n")
	var res [][]string
	for _, L := range lines { w := strings.Fields(L)
		w = processFlags(w)
		w = processPunctuation(w)
		w = processMarks(w)
		w = processArticle(w)
		res = append(res, w)}
return res}
