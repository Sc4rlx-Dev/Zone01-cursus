package main

import ("fmt" ;"os" ;"regexp";"strings")

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input file> <output file>")
		return
	}

	inputF := os.Args[1]
	outputF := os.Args[2]
	c, err := os.ReadFile(inputF)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if !strings.HasSuffix(inputF, ".txt") || !strings.HasSuffix(outputF, ".txt") {
		fmt.Println("Error Invalid format of file: should be file.txt\n<usage example: go run . sample.txt result.txt>")
		return
	}

	res := processText(string(c))
	final_res := ""

	for i, line := range res {
		final_res += strings.Join(line, " ")
		if i < len(res)-1 {
			final_res += "\n"
		}
	}

	rgx := regexp.MustCompile(`(\s+)([.,!?:;]+)`)
	final_res = rgx.ReplaceAllString(final_res, "${2}")
	rgx = regexp.MustCompile(`([.,!?:;]+)`)
	final_res = rgx.ReplaceAllString(final_res, "${1} ")
	rgx = regexp.MustCompile(`'\s*(.*?)\s*'`)
	final_res = rgx.ReplaceAllString(final_res, "'$1'")
	wordsfinal := strings.Fields(final_res)
	final_res = strings.Join(wordsfinal, " ")


	err = os.WriteFile(outputF, []byte(final_res), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Success!!")
}
