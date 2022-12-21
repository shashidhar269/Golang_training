package main

import (
	calculator "GO/Assignment02/calculator"
	helloworld "GO/Helloworld"

	"fmt"
)

func main() {
	// a, b := 4, 8
	var a, b int
	fmt.Printf("enter the values:")
	fmt.Scanf("%d,%d", &a, &b)
	addition := calculator.Add(a, b)
	// var x float64 = float64(addition)
	fmt.Println("addition of two numbers = ", addition)
	subtract := calculator.Sub(a, b)
	fmt.Println("subtract of two numbers = ", subtract)
	c, d := calculator.Div(a, b)
	var x float64 = float64(c)
	fmt.Printf("quotient = %f reminder = %d \n", x, d)
	multiplication := calculator.Mul(a, b)
	fmt.Println("multiplication of two numbers = ", multiplication)

	fmt.Println("This is from helloworld package = ", helloworld.Main())
}
