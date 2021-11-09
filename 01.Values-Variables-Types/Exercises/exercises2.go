package main

import "fmt"

func main () {
	pi := 3.14
	radius := 5

	circleCirCum := 2 * pi * float64(radius)
	fmt.Println(circleCirCum)

	fmt.Printf("For a radius of %v, the circle circumference is %.2f.", radius, circleCirCum)
}