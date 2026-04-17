package main

import (
	"strings"
)

func fixArticles(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word)-1; i++ {
		vow := strings.ContainsAny("aeiouhAEIOUH", string(word[i+1][:1]))
		if word[i] == "a" && vow {
			word[i] = "an"

		} else if word[i] == "A" && vow {
			word[i] = "An"
		} else if word[i] == "an" && !vow {
			word[i] = "a"
		} else if word[i] == "An" && !vow {
			word[i] = "A"
		}
	}
	return strings.Join(word, " ")

}
