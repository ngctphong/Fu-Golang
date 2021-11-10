package main

import "fmt"

var userName = "Max"

func main() {
	shouldContinue := true
	userName := "Maximilian" // shadow the higher-level variable

	if shouldContinue {
		userAge := 32

		fmt.Printf("Name: %v, Age: %v", userName, userAge)
	}

	fmt.Println()
	printData()
	fmt.Println()

	// new & make
	// hobbies := []string{"Sports", "Reading"}

	hobbies := make([]string, 2, 10)
	moreHobbies := new([]string)

	// aMap := make(map[string]int, 5)

	fmt.Println(hobbies)
	fmt.Println(*moreHobbies)

	hobbies[0] = "Sports"
	hobbies[1] = "Reading"

	hobbies = append(hobbies, "Cooking", "Dancing")

	fmt.Println(hobbies)

	*moreHobbies = append(*moreHobbies, "Sports")
	fmt.Println(*moreHobbies)
}

func printData() {
	fmt.Printf("Name: %v", userName)
}
