package main

import "fmt"

func main() {

	var i int = 0
	var soma int = 0
	for i <= 10 {
		soma += i
		i++
	}
	fmt.Println(soma)
}
