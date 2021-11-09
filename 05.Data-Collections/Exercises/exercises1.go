package main

import "fmt"

type Product struct {
	id string
	title string
	price float64
}

func main() {
	hobbies := [3]string{"game", "football", "code"}

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	newHobbies := hobbies[:2]
	newHobbies = newHobbies[1:]
	newHobbies = append(newHobbies, hobbies[2])
	fmt.Println(newHobbies)

	goals := []string{"learning", "doing"}
	goals[1] = "creating project"
	goals = append(goals, "earning")
	fmt.Println(goals)

	products := []Product{
		{
			"first-product",
			"A First Product",
			12.99,
		},
		{
			"second-product",
			"A Second Product",
			129.99,
		},
	}

	fmt.Println(products)

	newProduct := Product{
		"third-product",
		"A Third Product",
		15.99,
	}

	products = append(products, newProduct)
	fmt.Println(products)
}