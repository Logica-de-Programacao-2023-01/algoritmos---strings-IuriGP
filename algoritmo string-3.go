package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite a frase: ")
	scanner.Scan()
	s1 := scanner.Text()

	fmt.Print("Digite a letra: ")
	scanner.Scan()
	Lold := scanner.Text()

	fmt.Print("Digite a letra: ")
	scanner.Scan()
	Lnew := scanner.Text()

	newS := strings.ReplaceAll(s1, Lold, Lnew)
	fmt.Println(newS)
}
