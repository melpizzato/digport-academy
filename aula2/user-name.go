package main

import "fmt"

func main() {

	fmt.Println("Digite o seu nome:")

	var name string

	fmt.Scanln(&name)
	fmt.Printf("Bem-vindo %s", name)
}
