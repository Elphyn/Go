package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func userInput() (a, b int, operator string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите значения: ")
		input, _ := reader.ReadString('\n')
		splitInput := strings.Fields(input)
		fmt.Println(splitInput)

		a, _ := strconv.Atoi(splitInput[0])
		b, _ := strconv.Atoi(splitInput[2])
		//if splitInput[1] == "+" {
		//	result := a + b
		//	fmt.Printf("Add method chosen: %d\n", result)
		//}
		if !(a >= 0 && a <= 10 && b >= 0 && b <= 10) {
			fmt.Println("Error, input values must be in range from (0,10) including 0 and 10)")
			continue
		}
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

}
