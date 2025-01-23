package goreloded

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ProcessText(text string) string {
	lines := strings.Split(text, "\n")
	for i := range lines {
		lines[i] = applyMultiWordTransforms(lines[i])
		lines[i] = applySingleWordTransforms(lines[i])
		lines[i] = handlePunctuationSpacing(lines[i])
		lines[i] = handleEllipsesAndPunctuationGroups(lines[i])
		lines[i] = regexp.MustCompile(`'\s*(.*?)\s*'`).ReplaceAllString(lines[i], "'$1'")
		lines[i] = regexp.MustCompile(`\b(a)\s+([aeiouhAEIOUH])`).ReplaceAllString(lines[i], "an $2")
	}
	text = strings.Join(lines, "\n")
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")
	return strings.TrimSpace(text)
}

func applyMultiWordTransforms(text string) string {
	text = regexp.MustCompile(`(\b(?:\w+\b(?: \w+){0,})*) \(low, (\d+)\)`).ReplaceAllStringFunc(text, func(match string) string {
		return applyLowTransform(match)
	})

	text = regexp.MustCompile(`(\b(?:\w+\b(?: \w+){0,})*) \(up, (\d+)\)`).ReplaceAllStringFunc(text, func(match string) string {
		return applyUpTransform(match)
	})

	text = regexp.MustCompile(`(\b(?:\w+\b(?: \w+){0,})*) \(cap, (\d+)\)`).ReplaceAllStringFunc(text, func(match string) string {
		return applyCapTransform(match)
	})

	return text
}

func applyLowTransform(match string) string {
	parts := regexp.MustCompile(`(\b(?:\w+\b(?: \w+){0,})*) \(low, (\d+)\)`).FindStringSubmatch(match)
	words := strings.Fields(parts[1])
	n, _ := strconv.Atoi(parts[2])
	start := len(words) - n
	if start < 0 {
		start = 0
	}
	for i := start; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, " ")
}

func applyUpTransform(match string) string {
	parts := regexp.MustCompile(`(\b(?:\w+\b(?: \w+){0,})*) \(up, (\d+)\)`).FindStringSubmatch(match)
	words := strings.Fields(parts[1])
	n, _ := strconv.Atoi(parts[2])
	start := len(words) - n
	if start < 0 {
		start = 0
	}
	for i := start; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return strings.Join(words, " ")
}

func applyCapTransform(match string) string {
	parts := regexp.MustCompile(`(\b(?:\w+\b(?: \w+){0,})*) \(cap, (\d+)\)`).FindStringSubmatch(match)
	words := strings.Fields(parts[1])
	n, _ := strconv.Atoi(parts[2])
	start := len(words) - n
	if start < 0 {
		start = 0
	}
	for i := start; i < len(words); i++ {
		words[i] = strings.Title(strings.ToLower(words[i]))
	}
	return strings.Join(words, " ")
}

func applySingleWordTransforms(text string) string {
	transformations := map[*regexp.Regexp]func([]string) string{
		regexp.MustCompile(`(\b[0-9A-Fa-f]+\b) \(hex\)`): func(match []string) string {
			decimal, err := strconv.ParseInt(match[1], 16, 64)
			if err != nil {
				return match[1]
			}
			return fmt.Sprintf("%d", decimal)
		},
		regexp.MustCompile(`(\b[01]+\b) \(bin\)`): func(match []string) string {
			decimal, err := strconv.ParseInt(match[1], 2, 64)
			if err != nil {
				return match[1]
			}
			return fmt.Sprintf("%d", decimal)
		},
		regexp.MustCompile(`(\b\w+\b) \(up\)`): func(match []string) string {
			return strings.ToUpper(match[1])
		},
		regexp.MustCompile(`(\b\w+\b) \(low\)`): func(match []string) string {
			return strings.ToLower(match[1])
		},
		regexp.MustCompile(`(\b\w+\b) \(cap\)`): func(match []string) string {
			return strings.Title(strings.ToLower(match[1]))
		},
	}

	for reg, handler := range transformations {
		text = reg.ReplaceAllStringFunc(text, func(match string) string {
			parts := reg.FindStringSubmatch(match)
			return handler(parts)
		})
	}

	return text
}

func handlePunctuationSpacing(text string) string {
	return regexp.MustCompile(`\s*([.,!?;:])\s*`).ReplaceAllString(text, "$1 ")
}

func handleEllipsesAndPunctuationGroups(text string) string {
	return regexp.MustCompile(`\s*([.]{3}|[!?]{2,})\s*`).ReplaceAllStringFunc(text, strings.TrimSpace)
}
