package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findCommonDivisors(numbers []int) []int {
	// Начнем с нахождения наименьшего числа в массиве для того что бы понять вдруг он минимальный делитель
	minNumber := numbers[0]
	for _, number := range numbers {
		if number < minNumber {
			minNumber = number
		}
	}

	// Создадим слайс для общих делителей
	var commonDivisors []int

	// Переберем числа от 2 до minNumber. Можно начат перебор с числа 1 но он всегда будет в конечном выводе
	for divisor := 2; divisor <= minNumber; divisor++ {
		// Проверим, является ли divisor делителем всех чисел в массиве
		isCommonDivisor := true
		for _, number := range numbers {
			if number%divisor != 0 {
				isCommonDivisor = false
				break
			}
		}
		if isCommonDivisor {
			commonDivisors = append(commonDivisors, divisor)
		}
	}

	return commonDivisors
}

func main() {
	// Запросим у пользователя ввод чисел через запятую тут я использую буферизированный ввод он быстрее
	fmt.Print("Введите положительные целые числа: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Преобразуем строку в числа и запишем в массив
	inputNumbersStr := strings.Fields(input)
	var inputNumbers []int
	for _, numStr := range inputNumbersStr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Ошибка при преобразовании числа:", err)
			os.Exit(1)
		}
		inputNumbers = append(inputNumbers, num)
	}

	// Найдем общие делители
	result := findCommonDivisors(inputNumbers)
	fmt.Printf("Общие делители для чисел %v: %v\n", inputNumbers, result)
}
