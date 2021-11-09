package main

import "fmt"

func main() {
	age := 32
	fmt.Println(age)

	// myAge := &age
	var myAge *int
	myAge = &age
	fmt.Println(*myAge)

	*myAge = 33
	fmt.Println(age)

	doubleAge := double(myAge)
	fmt.Println(doubleAge)	
	fmt.Println(age)	
}

func double (number *int) int {
	result := *number * 2
	*number = 100
	return result
}