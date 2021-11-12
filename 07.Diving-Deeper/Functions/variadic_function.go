package main

import "fmt"

func main() {
	numbers := []int{1, 2, 5, 7, 8, 10}

	sum := sumup(24, 56, 78, 34)
	anotherSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
