package main

import (
	"GO/Assignment02/calculator"
	"fmt"
)

var z = [4]string{"hyd", "bang", "mysore", "chennai"}

func updateThirdElement(z *[4]string) *[4]string {
	z[2] = "Texas"
	return z
}

func main() {
	//Create an ARRAY which holds 5 VALUES of TYPE int
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Printf("Type of array = %T\n", a)
	fmt.Println("-----------------------")
	//Create a SLICE of TYPE int assign 10 VALUES
	b := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(b)
	fmt.Printf("Type of array = %T\n", b)
	fmt.Println("-----------------------")
	d := b[:5]
	fmt.Println(d)
	d = b[5:10]
	fmt.Println(d)
	d = b[2:7]
	fmt.Println(d)
	d = b[1:6]
	fmt.Println(d)
	fmt.Println("-----------------------")

	// 4
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	//fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
	//fmt.Println(cap(x))
	//fmt.Println(len(x))
	//fmt.Println(x[9])
	fmt.Println("-----------------------")
	fmt.Println("before", z)
	//updateThirdElement(&z)
	fmt.Println("after", updateThirdElement(&z))
	fmt.Println("-----------------------")

	//6
	var m, n int
	fmt.Printf("enter the values:")
	fmt.Scanf("%d,%d", &m, &n)
	addition := calculator.Add(&m, &n)
	// var x float64 = float64(addition)
	fmt.Println("addition of two numbers = ", addition)
	subtract := calculator.Sub(&m, &n)
	fmt.Println("subtract of two numbers = ", subtract)
	c, e := calculator.Div(&m, &n)
	fmt.Printf("quotient = %d reminder = %d \n", c, e)
	multiplication := calculator.Mul(&m, &n)
	fmt.Println("multiplication of two numbers = ", multiplication)

}
