package main

import (
	"fmt"
	
)

func arithmetics(x, y int) (result int) {
	fmt.Println("Addition: ", x+y)
	fmt.Println("Subtraction: ", x-y)
	fmt.Println("Multiplication: ", x*y)
	fmt.Println("Division: ", x/y)
	fmt.Println("Percentage: ", x%y)

	return
}

func relational(a, b int) {
	if a == b {
		fmt.Println("Both are equal")
	} else if a > b {
		fmt.Println(a, " is greater than ", b)
	} else {
		fmt.Println(a, " is less than ", b)
	}
}

func main() {
	arithmetics(20, 4)
	relational(10, 20)
}
