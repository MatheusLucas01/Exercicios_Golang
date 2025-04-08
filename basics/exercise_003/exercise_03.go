package main

import "fmt"

func carteiraHabilitacao() {
	var nome string
	var idade int

	fmt.Println("*Carteira Habilitação*")

	//Solicita o nome do usuário
	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)

	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&idade)

	if idade > 18 {
		fmt.Printf("Olá %s, você está pronto para tirar a carteira, já tem %d anos.", nome, idade)
	} else {
		fmt.Printf("Desculpe %s, você ainda é menor de idade.\nNão é possível agora. 😉", nome)
	}
}

func main() {
	carteiraHabilitacao()
}
