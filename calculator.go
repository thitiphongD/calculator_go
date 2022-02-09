package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		input = strings.TrimRight(input, "\r\n")
		message, err := fmt.Printf("%s ERR: not string must number only \n", input)
		if err != nil {
			fmt.Println(err)
		}
		panic(message)
	}

	return value

}

func addition(value1, value2 float64) float64 {
	return value1 + value2
}

func subtraction(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplication(value1, value2 float64) float64 {
	return value1 * value2
}

func division(value1, value2 float64) float64 {
	return value1 / value2
}

func getOperator() string {
	fmt.Print("Operator is ( + - * /):")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func main() {
	value1 := getInput(" value1 = ")
	value2 := getInput(" value2 = ")

	var result float64

	switch operator := getOperator(); operator {
	case "+":
		result = addition(value1, value2)
	case "-":
		result = subtraction(value1, value2)
	case "*":
		result = multiplication(value1, value2)
	case "/":
		result = division(value1, value2)

	default:
		panic("wrong operator")
	}
	fmt.Printf("result is %v", result)
}
