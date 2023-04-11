package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, a, b string
	fmt.Println("escreva uma string:")
	fmt.Scan(&x)
	fmt.Println("escreva uma letra para ser substituida:")
	fmt.Scan(&a)
	fmt.Println("escreva uma letra que sera substituir:")
	fmt.Scan(&b)
	resultado := strings.ReplaceAll(x, a, b)
	fmt.Println(resultado)
}
