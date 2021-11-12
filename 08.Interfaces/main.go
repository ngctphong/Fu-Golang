package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

type logger interface {
	log()
}

type logData struct {
	message  string
	fileName string
}

func (lData *logData) log() {
	err := ioutil.WriteFile(lData.fileName, []byte(lData.message), fs.FileMode(0644))

	if err != nil {
		fmt.Println("Storing the log data failed!")
	}
}

type loggableString string

func (text loggableString) log() {
	fmt.Println(text)
}

func main() {
	userLog := &logData{"User logged in", "user-log.txt"}
	// do more work ...
	createLog(userLog)

	message := loggableString("[INFO] User created!")
	// do more work ...
	createLog(message)

	// outputValue(userLog)
	// outputValue(message)

	firstNumber := 5
	secondNumber := 6.1
	numbers := []int{1, 2, 3}

	doubledFirstNumber := double(firstNumber)
	doubledSecondNumber := double(secondNumber)
	doubledNumbers := double(numbers)
	doubledString := double("Test")

	outputValue(doubledFirstNumber)
	outputValue(doubledSecondNumber)
	outputValue(doubledNumbers)
	outputValue(doubledString)

}

func createLog(data logger) {
	data.log()
}

func outputValue(value interface{}) {
	// val, ok := value.(loggableString)
	// fmt.Println(val, ok)
	fmt.Println(value)
}

func double(value interface{}) interface{} {
	switch val := value.(type) {
	case int:
		return val * 2
	case float64:
		return val * 2
	case []int:
		numNumbers := append(val, val...)
		return numNumbers
	default:
		return ""
	}
}
