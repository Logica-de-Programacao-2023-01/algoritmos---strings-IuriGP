package main

import (
	"fmt"
	"strings"
)

func main() {
	var frase string

	fmt.Print("Digite a frase: ")
	fmt.Scanln(&frase)

	result := strings.Builder{}
	vogais := "aeiouAEIOU"

	for _, char := range frase {
		if !strings.ContainsRune(vogais, char) {
			result.WriteRune(char)
		}
	}
	fmt.Println(result.String())
}
