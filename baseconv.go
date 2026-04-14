package main

import (
	"strconv"
	"strings"
)

func HexToDec(str string) string {
	word := strings.Fields(str)

	for i := 0; i < len(word); i++ {
		if word[i] == "(hex)" && i > 0 {
			value, err := strconv.ParseInt(word[i-1], 16, 64)
			if err == nil {
				word[i-1] = strconv.FormatInt(value, 10)
				word = append(word[:i], word[i+1:]...)
			}
		}

	}
	return strings.Join(word, " ")
}

func binToDec(str string) string {
	word := strings.Fields(str)

	for i := 0; i < len(word); i++ {
		if word[i] == "(bin)" && i > 0 {
			value, err := strconv.ParseInt(word[i-1], 2, 64)
			if err == nil {
				word[i-1] = strconv.FormatInt(value, 10)
				word = append(word[:i], word[i+1:]...)
			}
		}

	}
	return strings.Join(word, " ")
}
