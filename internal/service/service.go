package service

import (
	"strings"

	"github.com/qewyxiww/projectt/pkg/morse"
)

func DetectAndConvert(content string) (string, error) {
	if content == "" {
		return "", nil
	}

	content = strings.TrimSpace(content)

	if isMorseCode(content) {

		return morse.ToText(content), nil
	}
	return morse.ToMorse(content), nil
}

func isMorseCode(s string) bool {
	words := strings.Split(s, " / ")
	if len(words) == 0 {
		words = []string{s}
	}

	for _, word := range words {

		symbols := strings.Split(strings.TrimSpace(word), " ")

		for _, symbol := range symbols {
			symbol = strings.TrimSpace(symbol)
			if symbol == "" {
				continue
			}

			for _, char := range symbol {
				if char != '.' && char != '-' {
					return false
				}
			}
		}
	}

	return len(s) > 0
}
