package main

import "fmt"

type User struct {
	nome string
	idade int
}

func main() {

	//1 ex
	//user := User{
	//	nome: "Gabriel",
	//	idade: 18,
	//}

	//2 ex
	//user := User{
	//	"Gabriel",
	//	18,
	//}

	//3 ex
	var user User = User{
		"Gabriel",
		18,
	}

	fmt.Println(user)
}
