package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите значение: ")
	inputString, _ := reader.ReadString('\n')
	inputString = strings.TrimSpace(inputString)
	inputInt, _ := strconv.Atoi(inputString)
	fmt.Println(inputInt)
	// _ похоже для каких то ошибок, так как метод возвращает несколько аргументов
	// надо разобраться как считывать несколько переменных с линии в терминале
}
