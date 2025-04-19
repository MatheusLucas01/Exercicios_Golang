package main

import "fmt"

func verificaParOuImpar() {
	var numero int

	fmt.Println(" -- Bem vindo ao par ou impar -- ")
	fmt.Print("Digite um número para saber se é par ou impar: ")
	fmt.Scanln(&numero)

	if numero%2 == 0 {
		fmt.Printf("O número %d, é par!", numero)
	} else {
		fmt.Printf("O número %d, é ímpar!", numero)
	}
}

func main() {
	verificaParOuImpar()
}
