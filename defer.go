package main

import "fmt"

func main() {

	// defer, adiando execução da linha de código

	defer fmt.Println("Oi")
	fmt.Println("Oi Tudo bem?")
}