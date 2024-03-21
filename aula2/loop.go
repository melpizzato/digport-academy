package main

import "fmt"

func main() {
	contador := 10

	for contador != 0 {
		fmt.Printf("Contagem regressiva: %d\n", contador)
		contador -= 1
	}
	fmt.Println("Arrasou!")
}
