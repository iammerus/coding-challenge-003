package main

import (
	"strings"
)

func main() {
	rot13("test")
}

func rot13(input string) string {
	parts := strings.Split(input, "")

	for index, part := range parts {
		parts[index] = rotate(part, 13)
	}

	return strings.Join(parts, "")
}

func rotate(character string, number int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	if strings.ToUpper(character) == character {
		alphabet = strings.ToUpper(alphabet)
	}

	if strings.Index(alphabet, character) == -1 {
		return character
	}

	total := strings.Index(alphabet, character) + number

	var index = 0

	if total > 25 {
		index = total - 26
	} else {
		index = total
	}

	return string(alphabet[index])
}
