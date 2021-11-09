package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a, b := generateRandomNumber()

	sum := add(a, b)

	printNumber(sum)
}

func add(num1, num2 int) int {
	return num1 + num2
}

// func add(num1, num2 int) (sum int) {
// 	sum = num1 + num2
// 	return
// }

func printNumber(number int) {
	fmt.Println(number)
	fmt.Printf("Number: %v", number)
}

func generateRandomNumber() (int, int) {
	rand.Seed(time.Now().UnixNano())
	randomNumber1 := rand.Intn(10)
	randomNumber2 := rand.Intn(10)
	return randomNumber1, randomNumber2
}
