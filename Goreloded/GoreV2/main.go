package main

import (
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run . <input file> <output file>")
        return
    }

    inputF := os.Args[1]
    outputF := os.Args[2]

    // Read the input file
    content, err := os.ReadFile(inputF)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Process the text
    mtxt := ProcessText(string(content))

    // Write the processed text to the output file
    err = os.WriteFile(outputF, []byte(mtxt), 0644)
    if err != nil {
        fmt.Println("Error writing file:", err)
        return
    }
}

// ProcessText processes the input text according to the specified rules
func ProcessText(text string) string {
    // Define regex patterns for different modifications
    hexPattern := regexp.MustCompile(`(\b[0-9A-Fa-f]+\b) \(hex\)`)
    binPattern := regexp.MustCompile(`(\b[01]+\b) \(bin\)`)
    upPattern := regexp.MustCompile(`(\b\w+\b) \(up\)`)
    lowPattern := regexp.MustCompile(`(\b\w+\b) \(low\)`)
    capPattern := regexp.MustCompile(`(\b\w+\b) \(cap\)`)
    upNumPattern := regexp.MustCompile(`(\b\w+(?: \w+){0,}) \(up, (\d+)\)`)
    lowNumPattern := regexp.MustCompile(`(\b\w+(?: \w+){0,}) \(low, (\d+)\)`)
    capNumPattern := regexp.MustCompile(`(\b\w+(?: \w+){0,}) \(cap, (\d+)\)`)
    punctPattern := regexp.MustCompile(`\s*([.,!?;:])\s*`)
    quotePattern := regexp.MustCompile(`'\s*(.*?)\s*'`)
    aPattern := regexp.MustCompile(`\b(a)\s+([aeiouhAEIOUH])`)

    // Replace hex and bin patterns
    text = hexPattern.ReplaceAllStringFunc(text, func(s string) string {
        parts := hexPattern.FindStringSubmatch(s)
        hexValue, _ := strconv.ParseInt(parts[1], 16, 64)
        return fmt.Sprintf("%d", hexValue)
    })

    text = binPattern.ReplaceAllStringFunc(text, func(s string) string {
        parts := binPattern.FindStringSubmatch(s)
        binValue, _ := strconv.ParseInt(parts[1], 2, 64)
        return fmt.Sprintf("%d", binValue)
    })

    // Replace up, low, cap patterns
    text = upPattern.ReplaceAllStringFunc(text, func(s string) string {
        parts := upPattern.FindStringSubmatch(s)
        return strings.ToUpper(parts[1])
    })

    text = lowPattern.ReplaceAllStringFunc(text, func(s string) string {
        parts := lowPattern.FindStringSubmatch(s)
        return strings.ToLower(parts[1])
    })

    text = capPattern.ReplaceAllStringFunc(text, func(s string) string {
        parts := capPattern.FindStringSubmatch(s)
        return strings.Title(parts[1])
    })

    // Replace up, low, cap with numbers
    text = upNumPattern.ReplaceAllStringFunc(text, func(s string) string {
        parts := upNumPattern.FindStringSubmatch(s)
        num, _ := strconv.Atoi(parts[2])
        words := strings.Fields(parts[1])
        for i := 0; i < num && i < len(words); i++ {
            words[i] = strings.ToUpper(words[i])
        }
        return strings.Join(words, " ")
    })

    text = lowNumPattern.ReplaceAllStringFunc(text, func(s string) string {
        parts := lowNumPattern.FindStringSubmatch(s)
        num, _ := strconv.Atoi(parts[2])
        words := strings.Fields(parts[1])
        for i := 0; i < num && i < len(words); i++ {
            words[i] = strings.ToLower(words[i])
        }
        return strings.Join(words, " ")
    })

    text = capNumPattern.ReplaceAllStringFunc(text, func(s string) string {
        parts := capNumPattern.FindStringSubmatch(s)
        num, _ := strconv.Atoi(parts[2])
        words := strings.Fields(parts[1])
        for i := 0; i < num && i < len(words); i++ {
            words[i] = strings.Title(words[i])
        }
        return strings.Join(words, " ")
    })

    // Replace punctuation
    text = punctPattern.ReplaceAllStringFunc(text, func(s string) string {
        parts := punctPattern.FindStringSubmatch(s)
        return parts[1] // Return punctuation without spaces
    })

    // Handle groups of punctuation like "..."
    groupPunctPattern := regexp.MustCompile(`\s*([.]{3}|[!?]{2,})\s*`)
    text = groupPunctPattern.ReplaceAllStringFunc(text, func(s string) string {
        parts := strings.TrimSpace(s)
        return parts // Keep the group intact
    })

    // Add spaces after punctuation except for grouped ones
    singlePunctPattern := regexp.MustCompile(`\s*([.,!?;:])\s*`)
    text = singlePunctPattern.ReplaceAllString(text, "$1 ")

    text = strings.ReplaceAll(text, "  ", " ")
	// Punctuation tests are...kinda boring,what do you think?
	// Punctuation tests are... kinda boring, what do you think?


    // Replace quotes
    text = quotePattern.ReplaceAllString(text, "'$1'")

    // Replace 'a' with 'an' before vowels
    text = aPattern.ReplaceAllString(text, "an $2")

    return text
}