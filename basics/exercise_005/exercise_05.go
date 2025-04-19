package main

import "fmt"

// Programa que pede as 4 médias bimestrais e imprime a média
func media() {
	var media float64
	var nota1, nota2, nota3, nota4 float64
	var nome string

	fmt.Print("Olá, qual seu nome? ")
	fmt.Scanln(&nome)
	fmt.Println("Olá ", nome)

	fmt.Println("Digite a nota do 1° Bimestre: ")
	fmt.Scanln(&nota1)
	fmt.Println("Digite a nota do 2° Bimestre: ")
	fmt.Scanln(&nota2)
	fmt.Println("Digite a nota do 3° Bimestre: ")
	fmt.Scanln(&nota3)
	fmt.Println("Digite a nota do 4° Bimestre: ")
	fmt.Scanln(&nota4)

	media = (nota1 + nota2 + nota3 + nota4) / 4

	if media >= 6 {
		fmt.Printf("Parabéns, você passou, com nota máxima! Sua nota foi: %.2f ", media)
	} else {
		fmt.Println("Você reprovou! Sua nota foi:", media)
	}
}

func main() {
	media()
}
