package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite a frase: ")
	scanner.Scan()
	f1 := scanner.Text()

	fmt.Print("Digite a frase: ")
	scanner.Scan()
	f2 := scanner.Text()

	if f1 == f2 {
		fmt.Println("As strings são iguais")
	} else {
		fmt.Println("As strings são diferentes")
	}

}
