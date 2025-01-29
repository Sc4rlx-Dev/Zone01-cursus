package main

import ("strings")

func processMarks(words []string) []string {
	q := byte(39)

	ft_cm := func(slice []string) int { count := 0
		for _, word := range slice {
			count += strings.Count(word, string(q))}
		return count}
	for i := 0; i < len(words); {
		if words[i] == string(q) {

			if len(words) == 1 {break}

			if i == 0 && len(words) > 1 && !strings.HasPrefix(words[i+1], string(q)) { words[i+1] = string(q) + words[i+1]
				words = words[1:]

			} else if i == len(words)-1 && ft_cm(words[:i])%2 != 0 {
				words[i-1] += string(q)
				words = words[:i]

			} else if i > 0 && i < len(words)-1 {
				if ft_cm(words[:i])%2 != 0 { words[i-1] += string(q) } else if !strings.HasPrefix(words[i+1], string(q)) { words[i+1] = string(q) + words[i+1] }
				words = append(words[:i], words[i+1:]...)} else { i++ }
				continue}

		if strings.HasPrefix(words[i], string(q)) && len(words[i]) > 1 {
			if ft_cm(words[:i])%2 != 0 { words[i-1] += string(q)
				words[i] = words[i][1:] }}

		if strings.HasSuffix(words[i], string(q)) && len(words[i]) > 1 {
			if i < len(words)-1 && !strings.HasPrefix(words[i+1], string(q)) { words[i+1] = string(q) + words[i+1] 
				words[i] = words[i][:len(words[i])-1] }
		}
		i++}

	for i := range words {
		if len(words[i]) > 1 && words[i][0] == q && words[i][len(words[i])-1] == q {
			words[i] = strings.TrimSpace(words[i])}}

return words}
