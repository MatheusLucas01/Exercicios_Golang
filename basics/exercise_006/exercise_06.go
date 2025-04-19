package main

import "fmt"

func classificarIdade() {
	var idade int

	fmt.Println("** Bem vindo ao classificador de idade **")
	fmt.Println("Divirta-se!\n")

	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&idade)

	if idade <= 12 {
		fmt.Printf("Você tem %d anos, ainda é criança!", idade)
	} else if idade >= 13 && idade <= 17 {
		fmt.Printf("Você tem %d anos, já é um adolescente!", idade)
	} else if idade >= 18 && idade <= 59 {
		fmt.Printf("Você tem %d anos, é um adulto!", idade)
	} else {
		fmt.Printf("Você tem %d anos, é idoso!", idade)
	}
}

func main() {
	classificarIdade()
}
