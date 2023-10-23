package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func userInput() {
	reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Введите значение: ")
	//inputString, _ := reader.ReadString('\n')
	//inputString = strings.TrimSpace(inputString)
	//inputInt, _ := strconv.Atoi(inputString)
	//fmt.Println(inputInt)
	// _ похоже для каких то ошибок, так как метод возвращает несколько аргументов
	// надо разобраться как считывать несколько переменных с линии в терминале

	fmt.Println("Введите значения: ")
	input, _ := reader.ReadString('\n')
	splitInput := strings.Fields(input)
	fmt.Println(splitInput)
	//fmt.Printf("Type of splitInput: %s\n", reflect.TypeOf(splitInput))
	//fmt.Printf("Type of splitInput[0]: %s\n", reflect.TypeOf(splitInput[0]))
	a, _ := strconv.Atoi(splitInput[0])
	b, _ := strconv.Atoi(splitInput[2])
	if splitInput[1] == "+" {
		result := a + b
		fmt.Printf("Add method chosen: %d\n", result)
	}
}

func main() {
	userInput()
}
