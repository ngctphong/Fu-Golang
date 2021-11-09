package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ngctphong/bmi/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (float64, float64) {
	weight := getUserInput(info.WeightPrompt)
	height := getUserInput(info.HeightPrompt)

	return weight, height
}

func getUserInput(promptText string) float64 {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	enteredValue, _ := strconv.ParseFloat(userInput, 64)

	return enteredValue
}