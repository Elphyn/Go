package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// calculation part
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
		result := a / b
		return result
	}
	return 0
}

func isValidRoman(roman string) bool {
	romanNumerals := []string{
		"",
		"I",
		"II",
		"III",
		"IV",
		"V",
		"VI",
		"VII",
		"VIII",
		"IX",
		"X",
	}
	found := false
	for _, num := range romanNumerals {
		if num == roman {
			found = true
			break
		}
	}
	return found
}

func getIntFromRoman(roman string) int {
	reversedRomanNumerals := map[string]int{
		"I":     1,
		"II":    2,
		"III":   3,
		"IV":    4,
		"V":     5,
		"VI":    6,
		"VII":   7,
		"VIII":  8,
		"IX":    9,
		"X":     10,
		"XI":    11,
		"XII":   12,
		"XIII":  13,
		"XIV":   14,
		"XV":    15,
		"XVI":   16,
		"XVII":  17,
		"XVIII": 18,
	}
	value, _ := reversedRomanNumerals[roman]
	return value
}

func getRoman(number int) string {
	romanNumerals := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
		11: "XI",
		12: "XII",
		13: "XIII",
		14: "XIV",
		15: "XV",
		16: "XVI",
		17: "XVII",
		18: "XVIII",
	}
	value := romanNumerals[number]
	return value
}

// Handles User Input from console, checks conditions, returns 2 numbers and operator if conditions are met
func userInput() (first, second int, operator string) {
	//add roman bool to return, and roman input
	reader := bufio.NewReader(os.Stdin)
	for {
		//console input
		fmt.Println("Enter values: ")
		input, _ := reader.ReadString('\n')

		//cutting string to multiple parts
		splitInput := strings.Fields(input)

		//turning strings to numbers
		a, _ := strconv.Atoi(splitInput[0])
		b, _ := strconv.Atoi(splitInput[2])

		//return values: numbers
		first = a
		second = b
		//return value: operator (+, -, /, *)
		operator = splitInput[1]

		//checking task's conditions
		if !(a >= 0 && a <= 10 && b >= 0 && b <= 10) {
			fmt.Println("Error, input values must be in range from (0,10) including 0 and 10)")
			continue
		}
		//here exception triggers only if more than 3 arguments, but also user can type in number, operator, operator and so on, so condition need change
		if len(splitInput) > 3 {
			fmt.Println("Error, input values must be two numbers and one operator (+, -, /, *)")
			continue
		}
		break

	}
	return
}

func main() {

	//main loop
	//for {
	//	a, b, operator := userInput()
	//
	//	switch operator {
	//	case "+":
	//		fmt.Printf("Result: %d\n", Add(a, b))
	//	case "-":
	//		fmt.Printf("Result: %d\n", Sub(a, b))
	//	case "/":
	//		fmt.Printf("Result: %d\n", Div(a, b))
	//	case "*":
	//		fmt.Printf("Result: %d\n", Mul(a, b))
	//	}
	//
	//}

	n := 5
	fmt.Printf("%d in roman is: %s\n", n, getRoman(n))
	fmt.Printf("%s in decemal is: %d\n", getRoman(n), getIntFromRoman(getRoman(n)))
	fmt.Printf("%s is valid roman input?: %t", "XVIII", isValidRoman("XVIII"))
}

//Add roman numbers handling
//complete task's conditions, there's a lot of them
