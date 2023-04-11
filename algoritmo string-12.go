package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Println("escreva a string:")
	fmt.Scan(&x)

	letras := strings.Split(x, "")

	if len(letras)%2 != 0 {
		palindromo := true
		for i := 0; i < len(letras)/2; i++ {
			if letras[i] != letras[len(letras)-1-i] {
				palindromo = false
				break
			}
		}
		if palindromo {
			fmt.Print("palindromo")
		}
	} else {
		fmt.Printf("A string %s não é um palindromo")
	}
}
