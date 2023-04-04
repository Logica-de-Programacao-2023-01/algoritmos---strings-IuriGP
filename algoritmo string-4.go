package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite a primeira frase: ")
	scanner.Scan()
	s1 := scanner.Text()

	fmt.Print("Digite a segunda frase: ")
	scanner.Scan()
	s2 := scanner.Text()

	if s1 == s2 {
		fmt.Println("São frases identicas")
	} else {
		fmt.Println("São frases com algo diferente entre elas")
	}
}
