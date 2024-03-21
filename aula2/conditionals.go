package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

	if numero == 0 {
		fmt.Println("Zero!")
	} else if numero > 0 {
		fmt.Println("Positivo!")
	} else {
		fmt.Println("Negativo!")
	}
}
