package goreloded

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

)

// ProcessText processes the input text by applying various transformations based on specific patterns.
// The following transformations are supported:
// - Hexadecimal to decimal conversion: "<hex_value> (hex)" -> "<decimal_value>"
// - Binary to decimal conversion: "<binary_value> (bin)" -> "<decimal_value>"
// - Convert to uppercase: "<text> (up)" -> "<TEXT>"
// - Convert to lowercase: "<text> (low)" -> "<text>"
// - Capitalize each word: "<text> (cap)" -> "<Text>"
// - Convert first N words to uppercase: "<text> (up, <N>)" -> "<TEXT> <text>"
// - Convert first N words to lowercase: "<text> (low, <N>)" -> "<text> <text>"
// - Capitalize first N words: "<text> (cap, <N>)" -> "<Text> <text>"
// Additionally, it normalizes spaces around punctuation and handles specific text patterns:
// - Ensures single spacel around punctuation marks (.,!?;:)
// - Trims spaces around ellipses (...) and repeated punctuation (e.g., !!, ??)
// - Trims spaces inside single quotes
// - Replaces "a" with "an" before words starting with a vowel sound
// - Replaces double spaces with a single space
//
// Parameters:
// - text: The input string to be processed.
//
// Returns:
// - The processed string with all transformations applied.
func ProcessText(text string) string {
    patterns := map[*regexp.Regexp]func([]string) string{
        regexp.MustCompile(`(\b[0-9A-Fa-f]+\b) \(hex\)`): func(parts []string) string {
            hexValue, _ := strconv.ParseInt(parts[1], 16, 64)
            return fmt.Sprintf("%d", hexValue)
        },
        regexp.MustCompile(`(\b[01]+\b) \(bin\)`): func(parts []string) string {
            binValue, _ := strconv.ParseInt(parts[1], 2, 64)
            return fmt.Sprintf("%d", binValue)
        },
        regexp.MustCompile(`(\b\w+\b) \(up\)`): func(parts []string) string {
            return strings.ToUpper(parts[1])
        },
        regexp.MustCompile(`(\b\w+\b) \(low\)`): func(parts []string) string {
            return strings.ToLower(parts[1])
        },
        regexp.MustCompile(`(\b\w+\b) \(cap\)`): func(parts []string) string {
            return strings.Title(parts[1])
        },
        regexp.MustCompile(`(\b\w+(?: \w+){0,}) \(up, (\d+)\)`): func(parts []string) string {
            num, _ := strconv.Atoi(parts[2])
            words := strings.Fields(parts[1])
            for i := 0; i < num && i < len(words); i++ {
                words[i] = strings.ToUpper(words[i])
            }
            return strings.Join(words, " ")
        },
        regexp.MustCompile(`(\b\w+(?: \w+){0,}) \(low, (\d+)\)`): func(parts []string) string {
            num, _ := strconv.Atoi(parts[2])
            words := strings.Fields(parts[1])
            for i := 0; i < num && i < len(words); i++ {
                words[i] = strings.ToLower(words[i])
            }
            return strings.Join(words, " ")
        },
        regexp.MustCompile(`(\b\w+(?: \w+){0,}) \(cap, (\d+)\)`): func(parts []string) string {
            num, _ := strconv.Atoi(parts[2])
            words := strings.Fields(parts[1])
            for i := 0; i < num && i < len(words); i++ {
                words[i] = strings.Title(words[i])
            }
            return strings.Join(words, " ")
        },
    }

    for pattern, handler := range patterns {
        text = pattern.ReplaceAllStringFunc(text, func(s string) string {
            parts := pattern.FindStringSubmatch(s)
            return handler(parts)
        })
    }

    text = regexp.MustCompile(`\s*([.,!?;:])\s*`).ReplaceAllString(text, "$1 ")
    text = regexp.MustCompile(`\s*([.]{3}|[!?]{2,})\s*`).ReplaceAllStringFunc(text, strings.TrimSpace)
    text = regexp.MustCompile(`'\s*(.*?)\s*'`).ReplaceAllString(text, "'$1'")
    text = regexp.MustCompile(`\b(a)\s+([aeiouhAEIOUH])`).ReplaceAllString(text, "an $2")
    text = strings.ReplaceAll(text, "  ", " ")

	text = regexp.MustCompile(`\s+([.,!?;:])`).ReplaceAllString(text, "$1")
    return text
}
