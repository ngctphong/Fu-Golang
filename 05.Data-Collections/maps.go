package main

import "fmt"

func main() {
	websites := map[string]string {
		"Google": "https://google.com",
		"Amazon": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Amazon"])

	websites["Microsoft"] = "https://microsoft.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}