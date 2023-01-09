package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type person struct {
	firstname string
	lastname  string
	country   []string
}

type vehicle struct {
	doors int
	color string
}
type truck struct {
	fourWheel bool
	luxury    bool
	vehicle
}
type sedan struct {
	fourWheel bool
	luxury    bool
	vehicle
}

type shape interface {
	Area() float64
}

type square struct {
	l float64
	w float64
}

type circle struct {
	r float64
}

func (x square) Area() float64 {

	return x.l * x.w

}

func (x circle) Area() float64 {
	return math.Pi * x.r * x.r
}

func Info(s shape) {
	fmt.Println(s.Area())
}

func repetition(st string) map[string]int {

	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}

func main() {
	Person1 := person{
		firstname: "john",
		lastname:  "snow",
		country:   []string{"USA", "Australia", "Italy"},
	}
	Person2 := person{
		firstname: "Tyrion",
		lastname:  "Lannister",
		country:   []string{"Austria", "Turkey", "Dubai"},
	}
	fmt.Println("First person Struct ", Person1)
	fmt.Println("First person name ", Person1.firstname)
	fmt.Println("Looping over the favourite Country")
	for x, y := range Person1.country {
		fmt.Println(x, y)
	}
	fmt.Println("First person Struct ", Person2)
	fmt.Println("First person name ", Person2.firstname)
	fmt.Println(Person2.lastname)
	fmt.Println("Looping over the favourite Country")
	for x, y := range Person2.country {
		fmt.Println(x, y)
	}
	fmt.Println("-----------------------------------------------")
	//2
	personinfo := map[string]person{

		Person1.lastname: Person1,
		Person2.lastname: Person2,
	}
	for name, info := range personinfo {
		fmt.Printf("name:%s,firstname:%s,lastname:%s,Country:%s\n", name, info.firstname, info.lastname, info.country)
	}
	fmt.Println("-----------------------------------------------")
	//3
	veh1 := truck{true, false, vehicle{2, "black"}}
	fmt.Println(veh1)
	veh2 := sedan{false, true, vehicle{4, "white"}}
	fmt.Println(veh2)
	fmt.Println(veh1.vehicle.doors)
	fmt.Println(veh2.color)
	fmt.Println("-----------------------------------------------")
	//4
	x := square{
		l: 5.0,
		w: 10.0,
	}
	Info(x)
	y := circle{
		r: 10.0,
	}
	Info(y)

	fmt.Println("-----------------------------------------------")
	//5

	input := bufio.NewReader(os.Stdin)
	z, _ := input.ReadString('\n')
	z = strings.TrimSpace(z)
	//fmt.Println(z)
	for index, element := range repetition(z) {
		fmt.Println(index, "=", element)
	}
}
