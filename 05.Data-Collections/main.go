package main

import "fmt"

type Product struct {
	title string
	id string
	price float64
}

func main() {
	prices := []float64 {10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	productNames[2] = "A Carpet"

// 	fmt.Println(prices)
// 	fmt.Println(productNames)
// 	fmt.Println(prices[2])

// 	// featuredPrices := prices[1:3]
// 	// featuredPrices := prices[:2]
// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[:1]

// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(featuredPrices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))

// 	highlightedPrices = highlightedPrices[:3]

// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// }