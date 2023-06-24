package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Informe uma string: ")
	fmt.Scanln(&input)

	count := 0
	inWord := false

	for _, char := range input {
		if char == ' ' || char == '\t' || char == '\n' {
			inWord = false
		} else {
			if !inWord {
				count++
				inWord = true
			}
		}
	}
	fmt.Println("A string contém", count, "palavras.")
}
