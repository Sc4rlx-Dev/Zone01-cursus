package main

import ("fmt";"strconv";"strings")

func processFlags(words []string) []string {
	for i := 0; i < len(words); i++ {
		switch words[i] {
		
		case "(cap,", "(up,", "(low,":
			if i+1 < len(words) && strings.HasSuffix(words[i+1], ")") {
				n := strings.TrimSuffix(strings.TrimPrefix(words[i+1], ","), ")")
				c, err := strconv.Atoi(n)
				if err == nil {
					for j := 1; j <= c && i-j >= 0; j++ {
						switch words[i] {
						case "(cap,":
							words[i-j] = capitalize(words[i-j])
						case "(up,":
							words[i-j] = strings.ToUpper(words[i-j])
						case "(low,":
							words[i-j] = strings.ToLower(words[i-j])
						}
					}
					words = append(words[:i], words[i+2:]...)
					i--
				} else { fmt.Printf("error: invalid flag \"%s %s\"\n", words[i], words[i+1])
					words = append(words[:i], words[i+2:]...)
					return words }
			} else {fmt.Printf("error: invalid flag \"%s\"\n", words[i])
				words = append(words[:i], words[i+1:]...)
				return words}

		case "(cap)":
			if i > 0 {
				words[i-1] = capitalize(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			} else {
				words = append(words[:i], words[i+1:]...)
				i--}
		case "(up)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			} else {
				words = append(words[:i], words[i+1:]...)
				i--}
		case "(low)":
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			} else { words = append(words[:i], words[i+1:]...)
				i--}
		case "(hex)":
			if i > 0 {
				words[i-1] = hexToInt(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			} else { words = append(words[:i], words[i+1:]...)
				i--}
		case "(bin)":
			if i > 0 {
				words[i-1] = binToInt(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			} else { words = append(words[:i], words[i+1:]...)
				i--}
			}
	}
return words}
