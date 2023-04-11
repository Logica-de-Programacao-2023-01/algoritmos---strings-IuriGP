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

	_, err:=strconv.Atoi(x)
	if err != nil {
		fmt.Printf("A string %s não é uma sequencia")
	}else{
		numeros:=strings.Split(x,"")
		decresente:=true
		for i:=1;i<len(numeros);i++{

		}
	}
}
if decresente{
	fmt.println()
}