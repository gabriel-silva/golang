package main

import "fmt"

func main() {

	// 1 ex
	//var num int = 10
	//if num > 5 {
	//	fmt.Println("Número é maior que 5")
	//}

	//2 ex
	// inicialização no escopo do if
	if num := 10; num > 5 {
		fmt.Println("Número é maior que 5")
	}
}
