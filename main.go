package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Add(a, b int) int {
	result := a + b
	return result
}
func Sub(a, b int) int {
	result := a - b
	return result
}
func Mul(a, b int) int {
	result := a * b
	return result
}
func Div(a, b int) int {
	if a == 0 || b == 0 {
		fmt.Println("Error, division by 0 detected")
	} else {
		result := a * b
		return result
	}
	return 0
}

func userInput() (a, b int, operator string) {
	//add roman bool to return, and roman input, zero check
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите значения: ")
		input, _ := reader.ReadString('\n')
		splitInput := strings.Fields(input)
		fmt.Println(splitInput)

		//debug here shows that for some reason a, b = 0 here, need fix
		a, debug := strconv.Atoi(splitInput[0])
		b, _ := strconv.Atoi(splitInput[2])
		fmt.Println("Error: ", debug)
		//if splitInput[1] == "+" {
		//	result := a + b
		//	fmt.Printf("Add method chosen: %d\n", result)
		//}
		if !(a >= 0 && a <= 10 && b >= 0 && b <= 10) {
			fmt.Println("Error, input values must be in range from (0,10) including 0 and 10)")
			continue
		}
		//here exception triggers only if more than 3 arguments, but also user can type in number, operator, operator and so on, so condition need change
		if len(splitInput) > 3 {
			fmt.Println("Error, input values must be two numbers and one operator ()")
			continue
		}
		break

	}
	return
}

func main() {
	a, b, operator := userInput()
	fmt.Println(a)
	switch operator {
	case "+":
		fmt.Printf("Result: %d\n", Add(a, b))
	case "/":
		fmt.Printf("Result: %d\n", Div(a, b))

	}

}
