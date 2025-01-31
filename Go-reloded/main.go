package main

import ("fmt";"os";"strings")

func main() {
	if len(os.Args) != 3 {fmt.Println("Usage: go run . <input file> <output file>")
	return}

	inputF, outputF := os.Args[1], os.Args[2]
	c, err := os.ReadFile(inputF)

	if err != nil { fmt.Println("Error reading file:", err)
	return}

	if !strings.HasSuffix(inputF, ".txt") || !strings.HasSuffix(outputF, ".txt") { fmt.Println("Error Invalid format of file: should be file.txt\n<usage example: go run . sample.txt result.txt>")
	return}

	res := processText(string(c))
	wordProc := ""

	for i, line := range res {
		tmp := strings.Join(line, " ") + "\n"
		tmp = handlers(tmp)
		wordProc += tmp
		if i < len(res)-1 { wordProc += "\n" }
	}
	// final := handlers(wordProc)
	err = os.WriteFile(outputF, []byte(wordProc), 0644)
	if err != nil { fmt.Println("Error writing file:", err)
	return}

fmt.Println("Success!!")
}
