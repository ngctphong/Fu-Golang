package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func (p *Product) store() {
	file, _ := os.Create(p.id + ".txt")

	content := fmt.Sprintf(
		"ID: %v\nTitle: %v\nDescription: %v\nPrice: USD %.2f\n",
		p.id,
		p.title,
		p.description,
		p.price,
	)

	file.WriteString(content)
	file.Close()
}

func (p *Product) printData() {
	fmt.Printf("ID: %v\n", p.id)
	fmt.Printf("Title: %v\n", p.title)
	fmt.Printf("Description: %v\n", p.description)
	fmt.Printf("Price: %v\n", p.price)
	fmt.Println("---------------")
}

func NewProduct(id, t, d string, p float64) *Product {
	return &Product{id, t, d, p}
}

func main() {
	createdProduct := getProduct()

	createdProduct.printData()
	createdProduct.store()
}

func getProduct() *Product {
	fmt.Println("Please enter the product data.")
	fmt.Println("-----------------------------")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Product ID: ")

	idInput := readUserInput(reader, "Product ID: ")
	titleInput := readUserInput(reader, "Product Title: ")
	descriptionInput := readUserInput(reader, "Product Description: ")
	priceInput := readUserInput(reader, "Product Price: ")
	priceValue, _ := strconv.ParseFloat(priceInput, 64)

	product := NewProduct(idInput, titleInput, descriptionInput, priceValue)
	return product

}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}