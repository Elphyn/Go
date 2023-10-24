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
		result := a / b
		return result
	}
	return 0
}

func isValidRoman(roman string) bool {
	romanNumerals := []string{
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
		19: "XIX",
		20: "XX",
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
		"XIX":   19,
		"XX":    20,
	}
	value, _ := reversedRomanNumerals[roman]
	return value
}

func getRoman(number int) string {

	roman := ""
	// Define the values and corresponding symbols in parallel slices.
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(values); i++ {
		for number >= values[i] {
			roman += symbols[i]
			number -= values[i]
		}
	}

	return roman
}

// Handles User Input from console, checks conditions, returns 2 numbers and operator if conditions are met
func userInput() (first, second int, operator string, romanFlag bool) {
	reader := bufio.NewReader(os.Stdin)
	for {
		//romanFlag := false
		//console input
		fmt.Println("Enter values: ")
		input, _ := reader.ReadString('\n')

		//cutting string to multiple parts
		splitInput := strings.Fields(input)

		if isValidRoman(splitInput[0]) || isValidRoman(splitInput[2]) {
			romanFlag = true
			if !(isValidRoman(splitInput[0]) && isValidRoman(splitInput[2])) {
				fmt.Println("If you intend to use Roman, both values should be Roman")
				continue
			}
			a := getIntFromRoman(splitInput[0])
			b := getIntFromRoman(splitInput[2])

			//return values: numbers
			first = a
			second = b
			//return value: operator (+, -, /, *)
			operator = splitInput[1]

			possibleOperators := []string{"+", "-", "/", "*"}

			found := false
			for _, op := range possibleOperators {
				if op == operator {
					found = true
					break
				}
			}
			if !(found) {
				fmt.Println("Error, incorrect format: input values must be two numbers and one operator (+, -, /, *)")
				continue
			}
			if len(splitInput) > 3 {
				fmt.Println("Error, incorrect format: input values must be two numbers and one operator (+, -, /, *)")
				continue
			}
			if !(a >= 0 && a <= 10 && b >= 0 && b <= 10) {
				fmt.Println("Error, input values must be in range from (0,10) including 0 and 10)")
				continue
			}
			break
		} else {
			//turning strings to numbers
			a, exc_1 := strconv.Atoi(splitInput[0])
			b, exc_2 := strconv.Atoi(splitInput[2])

			//return values: numbers
			first = a
			second = b
			//return value: operator (+, -, /, *)
			operator = splitInput[1]

			possibleOperators := []string{"+", "-", "/", "*"}

			found := false
			for _, op := range possibleOperators {
				if op == operator {
					found = true
					break
				}
			}
			if !(found) {
				fmt.Println("Error, incorrect format: input values must be two numbers and one operator (+, -, /, *)")
				continue
			}
			//checking task's conditions
			if (exc_1 != nil || exc_2 != nil) || len(splitInput) > 3 {
				fmt.Println("Error, incorrect format: input values must be two numbers and one operator (+, -, /, *)")
				continue
			}
			if !(a >= 0 && a <= 10 && b >= 0 && b <= 10) {
				fmt.Println("Error, input values must be in range from (0,10) including 0 and 10)")
				continue
			}
			break
		}

	}
	return
}

func main() {

	//main loop

	for {

		a, b, operator, roman := userInput()

		switch operator {
		case "+":
			if roman {
				fmt.Printf("Result: %s\n", getRoman(Add(a, b)))
			} else {
				fmt.Printf("Result: %d\n", Add(a, b))
			}
		case "-":
			if roman {

				result := Sub(a, b)
				if result < 1 {
					fmt.Println("Error, Result with Roman numbers could only be positive")
				} else {
					fmt.Printf("Result: %s\n", getRoman(result))
				}

			} else {
				fmt.Printf("Result: %d\n", Sub(a, b))
			}
		case "/":
			if roman {
				fmt.Printf("Result: %s\n", getRoman(Div(a, b)))
			} else {
				fmt.Printf("Result: %d\n", Div(a, b))
			}
		case "*":
			if roman {
				fmt.Printf("Result: %s\n", getRoman(Mul(a, b)))
			} else {
				fmt.Printf("Result: %d\n", Mul(a, b))
			}
		}

	}

}
