package main

import "fmt"

func main() {
	var notas map[string]int
	notas = make(map[string]int)

	// Ana --> 9
	// Gabriel --> 10

	notas["Ana"] = 9
	notas["Gabriel"] = 10
	fmt.Println(notas)

	// value = "João" / isExist = true/false
	value, isExist := notas["João"]
	if isExist {
		fmt.Println(value)
	}

}
