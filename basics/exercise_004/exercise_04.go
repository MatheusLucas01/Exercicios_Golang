package main

import "fmt"

func metrosCentimetros() {
	var metros float64

	fmt.Printf("Digite o tamanho em metros para saber os centímetros: ")
	fmt.Scanln(&metros)

	centimetros := metros * 100
	fmt.Printf("O tamanho em centimetros é %.2f\n", centimetros)

}

func main() {
	metrosCentimetros()
}
