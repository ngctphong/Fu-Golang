
package main

import (
	"fmt"
	"github.com/ngctphong/firstapp/greeting"
)

func main() {
	// var greetingText string
	// greetingText = "Hello World!!!!!"

	// var greetingText string = "Hello World!!!!!"

	luckyNumber := 17
	evenMoreLuckyNumber := luckyNumber + 5

	fmt.Println(greeting.GreetingText)
	fmt.Println(greeting.GreetingText)
	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)

	var newNumber float64 = float64(luckyNumber) / 3

	fmt.Println(newNumber)

	var defaultFloat float64 = 1.23456789123456789123456
	var smallFloat float32 = 1.23456789123456789123456

	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)

	var firstRune rune = '$'
	fmt.Println(firstRune) // 36
	fmt.Println(string(firstRune)) // $

	var firstByte byte = 'a'
	fmt.Println(firstByte) // 97
	fmt.Println(string(firstByte)) // a

	firstName := "Phong"
	lastName := "Nguyen"
	fmt.Println(firstName + " " + lastName)

	fullName := firstName + " " + lastName
	age := 32

	fmt.Printf("Hi, I am %v and I am %v (Type: %T) years old.", fullName, age, age)
	fmt.Println()

	fullName = fmt.Sprintf("%v %v", firstName, lastName)
	fmt.Printf("Hi, I am %v and I am %v (Type: %T) years old.", fullName, age, age)
	fmt.Println()

	const pi = 3.14
	fmt.Println(pi)
}