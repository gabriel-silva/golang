package main

import "fmt"

// 1 ex
//func soma(x, y int) int {
//	return x + y
//}

// 2 ex
func soma(x int, y int) int {
	return x + y
}

//// 3 ex
//func calcular(x int) (int, int) {
//	var quadrado int =  x * x
//	var cubo int = x * x * x
//	return quadrado, cubo
//}

// 4 ex
func calcular(x int) (quadrado int, cubo int) {
	quadrado =  x * x
	cubo = x * x * x
	return
}

func main() {
	fmt.Println("soma:", soma(3, 4))
	fmt.Println(calcular(2))
}
