package main

import "fmt"

func Tabuada() {
	var numero int

	fmt.Print("Digite um número para ver sua tabuada: ")
	fmt.Scanln(&numero)

	fmt.Printf("\nTabuada do %d:\n", numero)
	for i := 1; i <= 10; i++ {
		resultado := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, resultado)
	}
}

func main() {
	Tabuada()
}

// for i := 1 --> Declara e inicializa com 1
// i <= 10 --> Vai de 1 até 10
// i++ --> Soma i + 1
