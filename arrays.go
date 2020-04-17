package main

import "fmt"

func main() {
	//1 ex
	//var numbers[5] int
	//numbers[0] = 1
	//numbers[1] = 2
	//numbers[2] = 3
	//numbers[3] = 4
	//numbers[4] = 5

	// 2 ex
	var numbers = [5] int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// array with for
	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i])
	}
}
