package main

import "fmt"

type tranformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{2, 3, 4}

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	tranformFn1 := getTransformerFunction(&numbers)
	tranformFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers1 := transformNumbers(&numbers, tranformFn1)
	transformedNumbers2 := transformNumbers(&moreNumbers, tranformFn2)

	fmt.Println(transformedNumbers1)
	fmt.Println(transformedNumbers2)
}

func transformNumbers(numbers *[]int, transform tranformFn) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformerFunction(numbers *[]int) tranformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}
