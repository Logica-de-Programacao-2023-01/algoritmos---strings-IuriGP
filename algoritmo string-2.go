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
	newS := strings.ReplaceAll(s1, " ", "_")
	fmt.Println(newS)
}
