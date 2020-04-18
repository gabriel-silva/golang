package main

import "fmt"

func main() {
	var numbers = [5] int{1, 2, 3, 4, 5}

	// utilizando range mas sem utilizar o index(i)
	//for _, value := range numbers {
	//	fmt.Println(value)
	//}

	// utilizando range com o index(i)
	for i, value := range numbers {
		fmt.Printf("index %d, value %d\n", i, value)
	}
}