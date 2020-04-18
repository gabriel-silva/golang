package main

import "fmt"

func main() {
	var numbers = [5] int{1, 2, 3, 4, 5}

	var slice []int = make([]int, 2)
	copy(slice, numbers[0:3])
	fmt.Println("slice:", slice)

	numbers[0] = 3
	//Slice
	fmt.Println("numbers: ", numbers)

	fmt.Println("slice2:", slice)
}
