package main

import "fmt"

func welcome() {
	var nome string
	var idade int

	// Solicita o nome do usuário
	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)

	// Solicita a idade do usuário
	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&idade)

	// Exibe a saudação formatada
	// %s formata string, %d formata números inteiros
	fmt.Printf("Olá, %s! Você tem %d anos.\n", nome, idade)
}

func main() {
	welcome()
}
