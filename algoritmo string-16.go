package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, y string
	fmt.Println("escreva a string:")
	fmt.Scan(&x)
	fmt.Println("escreva a string:")
	fmt.Scan(&y)

	if strings.Contains(x, y) {
		fmt.Println("A string está contida")
	} else {
		fmt.Println("A strings não esta contida")
	}
}
