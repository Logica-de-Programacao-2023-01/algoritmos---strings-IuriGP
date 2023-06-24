package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Informe uma string: ")
	fmt.Scanln(&input)

	reversed := ""
	for _, char := range input {
		reversed = string(char) + reversed
	}

	fmt.Println("String invertida:", reversed)
}
