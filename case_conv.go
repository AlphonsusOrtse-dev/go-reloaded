package main

import (
    "fmt"
    "strings"
)

func conv(str string) string {
    word := strings.Fields(str)

    for v := 0; v < len(word); v++ {
        switch word[v] {
        case "(up)":
            word[v-1] = strings.ToUpper(word[v-1])
			word = append(word[:v], word[v+1:]...)

			
			
		}
	}
	result := strings.Join(word, " ")
	result = strings.ReplaceAll(result, " !", "!")
	return result

}

func main() {
	fmt.Println(conv("Ready, set, go (up) !" ))

}