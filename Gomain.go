package main

import (
	"fmt"
	"goreloded"
	"os"
)

func main() {    
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run . <input file> <output file>")
        return
    }
    
    inputF := os.Args[1]
    outputF := os.Args[2]

    content, err := os.ReadFile(inputF)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    processedText := goreloded.ProcessText(string(content))

    err = os.WriteFile(outputF, []byte(processedText), 0644)
    if err != nil {
        fmt.Println("Error writing file:", err)
        return
    }
	fmt.Println("Done")
}
