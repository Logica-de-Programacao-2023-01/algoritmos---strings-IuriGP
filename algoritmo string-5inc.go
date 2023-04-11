package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x string
	fmt.Println("escreva a string:")
	fmt.Scan(&x)

	numero, err := strconv.ParseFloat(x, 64)
	if err != nil {
		fmt.Printf("A string %s nao é flutuante", x)
	} else {
		fmt.Printf("O ponto flutuante é %s", x)
	}
}
