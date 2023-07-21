package main

import "fmt"

type Calculation struct {
	int1 float64
	int2 float64
}

func main() {
	calc := Calculation{
		int1: 32.3,
		int2: 23.2,
	}
	fmt.Println("Sum is ", calc.int1+calc.int2)
	fmt.Println("Diff is ", calc.int1-calc.int2)
	fmt.Println("Pdt is ", calc.int1*calc.int2)
	fmt.Println("Div is ", calc.int1/calc.int2)

	//pointers
	var a *int
	x := 56
	a = &x
	fmt.Println("The actual value: ", x)
	fmt.Println("The memory address of the value: ", a)
	fmt.Println("The value of the address: ", *a)
}
