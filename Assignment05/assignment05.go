package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Print(i, " ")

	}
	fmt.Print("\n")
	var i int = 1
	for i <= 50 {
		if i%2 == 1 {
			fmt.Print(i, " ")

		}
		i++
	}
	fmt.Print("\n")
	i = 1
	for {
		if i <= 50 {
			i += 1
			fmt.Print(i, " ")
			i++
		}
		if i >= 50 {
			break
		}

	}
	fmt.Print("\n")
	for i := 50.0; i <= 105.0; i++ {
		sum := i / 6.0
		fmt.Println("quotient of ", i, "=", sum)
	}

	//var y string
	//fmt.Print("Enter the string:")
	//fmt.Scanf("%s", &y)
	//fmt.Println(y)
	fmt.Println("-----------------")
	x := bufio.NewReader(os.Stdin)
	y, _ := x.ReadString('\n')
	y = strings.TrimSpace(y)
	fmt.Println(y)
	if y == "Golang tutorial" {
		fmt.Println("welcome")
	} else {
		fmt.Println("end")
	}
	for i := 1; i <= 80; i++ {
		if i%2 == 0 && i%4 == 0 {
			fmt.Println(i, "Golang tutorial")
		} else if i%2 == 0 && i%4 != 0 {
			fmt.Println(i, "Golang")
		} else if i%4 == 0 {
			fmt.Println(i, "tutorial")
		} else {
			fmt.Println(i)
		}
	}
}
