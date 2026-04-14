package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("go run . sample.txt result.txt")
		os.Exit(1)
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	read, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}
	data := string(read)
	result := conv(data)
	result = HexToDec(result)
	result = binToDec(result)

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("error writing file")
		os.Exit(1)

	}

}