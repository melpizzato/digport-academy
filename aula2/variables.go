package main

import "fmt"

func main() {

	num1 := 2
	var num2 int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num2)

	sum := num1 + num2

	fmt.Printf("Resultado: %d", sum)

}
