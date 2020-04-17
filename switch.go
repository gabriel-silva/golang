package main

import "fmt"

func main() {
	var nome string = "João"

	//1 ex
	//switch nome {
	//case "Ana":
	//	fmt.Println("é a Ana")
	//case "João":
	//	fmt.Println("é o João")
	//default:
	//	fmt.Println("Nada foi encontrado")
	//}

	//2 ex
	switch true {
	case nome == "Ana":
		fmt.Println("é a Ana")
	case nome == "João":
		fmt.Println("é o João")
	default:
		fmt.Println("Nada foi encontrado")
	}
}
