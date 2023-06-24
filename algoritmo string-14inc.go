package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var x string
	fmt.Println("escreva a string:")
	fmt.Scan(&x)

	x = strings.TrimSpace(x)
	numeros := strings.Fields(x)

	decrescente := true
	for i := 0; i < len(numeros)-1; i++ {
		numeroAtual, _ := strconv.Atoi(numeros[i])
		numeroProximo, _ := strconv.Atoi(numeros[i+1])

		if numeroAtual <= numeroProximo {
			decrescente = false
		}
	}
	if decrescente == true {
		fmt.Println("A sequência é numérica decrescente.")
	} else {
		fmt.Println("A sequência não é numérica decrescente.")
	}
}
